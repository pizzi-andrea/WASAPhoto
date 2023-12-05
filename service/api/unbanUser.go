package api

import (
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
taken uid of the user who wants unbband id of user banned and delete last one
*/
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var from, to int
	var err error
	var tk *security.Token
	var isBan bool

	/*
		Parse URL parameters
	*/

	// get :uid parameter in path
	if from, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	// get :bannedId parameter in path
	if to, err = strconv.Atoi(ps.ByName("bannedId")); err != nil {
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	/*
		if user id in URL path not exist, then user not found
	*/
	if isBan, err = rt.db.IsBanned(database.Id(from), database.Id(to)); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
	}

	if !isBan {
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Not Found, User not found")
	}

	/*
		Secure Bearer Authentication
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") // 401
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	/*
		Check if user can would put ban is account owner
	*/
	if tk.Value != database.Id(from) {
		w.Header().Set("content-type", "text/plain") // 403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		io.WriteString(w, UnauthorizedToken.Status)
		return

	}

	/*
		Try to remove ban
	*/

	if _, err = rt.db.DelBan(database.Id(from), database.Id(to)); err != nil {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
	}

	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusNoContent)

}
