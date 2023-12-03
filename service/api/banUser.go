package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var result bool
	var from, to int
	var err error
	var tk *security.Token

	/*
		Parse URL parameters
	*/
	if from, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	if to, err = strconv.Atoi(ps.ByName("bannedId")); err != nil {
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	/*
		if user id in URL path not exist, then user not found
	*/
	if _, err = rt.db.GetUserFromId(database.Id(from)); err != nil {
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Not found, user not found")
		return
	}

	if _, err = rt.db.GetUserFromId(database.Id(to)); err != nil {
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Not found, user not found")
		return
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
		Check if user  that want ban user can it. Only user have put ban
		can delete it
	*/
	if tk.Value != database.Id(from) {
		w.Header().Set("content-type", "text/plain") // 403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		io.WriteString(w, UnauthorizedToken.Status)
		return

	}

	/*
		Try to ban user
	*/

	if result, err = rt.db.PutBan(database.Id(from), database.Id(to)); err != nil {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
	}

	if result {
		var u *database.User
		if u, err = rt.db.GetUserFromId(database.Id(to)); u != nil || err == nil {
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(ServerError.StatusCode)
			io.WriteString(w, ServerError.Status)
			return
		}
		/* TODO: check error */
		rt.db.DelFollow(database.Id(from), database.Id(to)) // if user put ban auto. lost the follow
		rt.db.DelFollow(database.Id(to), database.Id(from)) // if user recive ban auto. lost the follow
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(*u)

	} else {
		w.Header().Set("content-type", "text/plain") //204
		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, "empty response, user just banned")
	}
}
