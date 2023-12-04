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
		gived uid and *followedId* then remove follower *followerId* from user followers
	      security:
*/
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var from, to int
	var err error
	var tk *security.Token
	var isBan bool
	var action bool
	var user *database.User
	/*
		Parse URL parameters in path
	*/
	if from, err = strconv.Atoi(ps.ByName("followerId")); err != nil {
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	if to, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	/*
		if user id in URL path not exist, then user not found
	*/

	if user, err = rt.db.GetUserFromId(database.Id(to)); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	if user == nil {
		fmt.Println(fmt.Errorf("not found %w", err))
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Not Found, User not found")
		return
	}

	/*
		Check if user that wont put follows can do it
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") // 401
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	if tk.Value != uint64(from) {
		w.Header().Set("content-type", "text/plain") // 403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		io.WriteString(w, UnauthorizedToken.Status)
		return

	}

	/*
		get banned user and check if not banned
	*/

	if isBan, err = rt.db.IsBanned(database.Id(to), database.Id(from)); isBan {
		w.Header().Set("content-type", "text/plain") // 403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		io.WriteString(w, UnauthorizedToken.Status)
		return

	}
	if err != nil {
		fmt.Println(fmt.Errorf("internal error: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, ServerError.Status)
		return
	}

	if action, err = rt.db.PutFollow(database.Id(from), database.Id(to)); err != nil {
		w.Header().Set("content-type", "text/plain") // 204
		w.WriteHeader(http.StatusNoContent)
		return

	}

	if action {

		var u *database.User
		if u, err = rt.db.GetUserFromId(database.Id(to)); u != nil || err == nil {
			w.Header().Set("content-type", "text/plain") // 500
			w.WriteHeader(ServerError.StatusCode)
			io.WriteString(w, ServerError.Status)
			return
		}
		w.Header().Set("content-type", "application/json") //201
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(*u)
	} else {
		w.Header().Set("content-type", "text/plain") //204
		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, "Empty response, just follow the userd")
	}
}
