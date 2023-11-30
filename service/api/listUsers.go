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

/*
this method return the list of users currently registered
*/

func (rt *_router) listUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var limit int
	var err error
	var users []database.User
	var tk *security.Token

	// get values from URL

	pLimit := r.URL.Query().Get("limit")
	pUsername := r.URL.Query().Get("username")

	/*
		Applay barrear authentication. if the user has passed the login phase then he will be able to access
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "plain/text")
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	/*parse URL parameters */
	if limit, err = strconv.Atoi(pLimit); pLimit != "" && err != nil {
		w.Header().Set("content-type", "plain/text") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return

	}

	if users, err = rt.db.GetUsers(pUsername); pUsername != "" && err != err {
		w.Header().Set("content-type", "plain/text") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	//future dev-----------------------------------------
	offset := 0
	if offset > len(users) {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}
	if limit == 0 {
		limit = len(users)
	}
	//----------------------------------------------------

	// response object
	users = users[offset:min(len(users), limit)]

	// if response body is empty
	if len(users) == 0 {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, "empty body, found nothing") // 204
		return
	}

	// write values getted in to response
	w.Header().Set("content-type", "application/json") // 200
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}
