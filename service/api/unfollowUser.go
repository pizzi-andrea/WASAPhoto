package api

import (
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var uid int
	var err error
	var uidUnfoll int
	var isFollower bool
	var tk *security.Token
	/*
		Parse URL parameters
	*/
	if uid, err = strconv.Atoi(ps.ByName("followerId")); err != nil {
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	if uidUnfoll, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	/*
		if folow not exist path URL path not exist
	*/

	if isFollower, err = rt.db.IsFollower(database.Id(uid), database.Id(uidUnfoll)); err != nil {
		w.Header().Set("content-type", "text/plain") // 500 code
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return

	}

	if !isFollower {
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Not found, just no follow user or user Id not exist")
		return

	}

	/*
		Applay barrear authentication. only owner can deleted follow
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") // 401
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	/*
		checks if the user who wants make unfollow is owner
	*/
	if tk.Value != uint64(uid) {
		w.Header().Set("content-type", "text/plain") // 403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		io.WriteString(w, UnauthorizedToken.Status)
		return
	}

	if _, err = rt.db.DelFollow(database.Id(uid), uint64(uidUnfoll)); err != nil {
		w.Header().Set("content-type", "text/plain") // 500 code
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	w.Header().Set("content-type", "text/plain") // 204
	w.WriteHeader(http.StatusNoContent)

}
