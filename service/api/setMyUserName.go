package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
assign username or update old username with new username.
The username to set is in the body request
*/

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var uid int
	var u string
	var err error
	var user *database.User

	var tk *security.Token

	/*
		Parse URL parameters
	*/
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	/*
		if user id in URL path not exist, then user not found
	*/
	if user, err = rt.db.GetUserFromId(database.Id(uid)); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	if user == nil {
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "user not found")
		return

	}

	/*
		Decode values in body request *r
	*/
	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		println(fmt.Errorf("%w", err))
		w.Header().Set("content-type", "text/plain") //400
		w.WriteHeader(BadRequest.Request.Response.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return

	}

	/*
		Applay barrear authentication. Username can update only his username
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") // 401
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	/*
		checks if the user who wants to change username is the owner
	*/
	if tk.Value != uint64(uid) {
		w.Header().Set("content-type", "text/plain") // 403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		io.WriteString(w, UnauthorizedToken.Status)
		return
	}

	/*
		Update username
	*/
	if _, err = rt.db.SetUsername(database.Id(uid), u); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusNoContent)
	io.WriteString(w, "Success, assigned new username")

}
