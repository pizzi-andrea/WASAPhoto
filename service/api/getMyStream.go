package api

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
given *uid* of user that who wants to get all photo in his stream
*/
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var uid int
	var err error
	var stream database.StreamPhotos
	var user *database.User
	var tk *security.Token
	var limit int
	offset := 0

	/*
		Parse URL parameters
	*/
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	pLimit := r.URL.Query().Get("limit")
	pUsername := r.URL.Query().Get("username")

	if limit, err = strconv.Atoi(pLimit); pLimit != "" && err != nil {
		fmt.Println(fmt.Errorf("atoi error: %w", err))
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return

	}

	rr, err := regexp.MatchString("^.*?$", pUsername)
	if !(rr && err == nil && len(pUsername) <= 16) {
		fmt.Println("username format error")
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
	}

	if user == nil {
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "user not found")
		return
	}

	/*

	 */
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") // 401
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	/*
		The stream is personal only owner user can see it
	*/
	if tk.Value != user.Uid {
		w.Header().Set("content-type", "text/plain") // 403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		io.WriteString(w, UnauthorizedToken.Status)
		return
	}

	if stream, err = rt.db.GetMyStream(user.Uid, pUsername, true, []database.OrderBy{}); err != nil {
		fmt.Println("query error: no offset valid")
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return

	}

	if limit == 0 || limit > len(stream) {
		limit = len(stream)
	}

	if offset > len(stream) { //500 response
		fmt.Println("offser: no offset valid")
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	stream = stream[offset:limit]

	if len(stream) == 0 { // 204 response
		fmt.Println("response: empty")
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, "empty body, found nothing")
		return

	} else {
		_ = stream /*TODO: misssing response for photo stream*/
	}

}
