package api

import (
	"encoding/json"
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
	// validate username in query
	rr, err := regexp.MatchString("^.*?$", pUsername)
	if !(len(pUsername) >= 1 && rr && err == nil && len(pUsername) <= 16) {
		fmt.Println("username format error")
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	/*
		Applay barrear authentication. if the user has passed the login phase then he will be able to access
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		fmt.Println(fmt.Errorf("authentication error: %w", err))
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	/*parse URL parameters */
	if limit, err = strconv.Atoi(pLimit); pLimit != "" && err != nil {
		fmt.Println(fmt.Errorf("atoi error: %w", err))
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return

	}

	if users, err = rt.db.GetUsers(pUsername, true); err != nil {
		fmt.Println(fmt.Errorf("query getUsers error: %w", err))
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	//future dev-----------------------------------------
	offset := 0
	_ = offset
	//-----------------------------------------------------------
	// response object

	if limit == 0 || limit > len(users) {
		limit = len(users)
	}
	var usersOk []database.User
	users = users[offset:limit]
	// if response body is empty
	if len(users) == 0 {
		fmt.Println("empty body")
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusNoContent)

		return
	}
	var b bool

	for _, u := range users {
		if b, err = rt.db.IsBanned(u.Uid, tk.Value); err != nil {
			fmt.Println(fmt.Errorf("query ban error: %w", err))
			w.Header().Set("content-type", "text/plain") // 500
			w.WriteHeader(ServerError.StatusCode)
			io.WriteString(w, ServerError.Status)
			return
		}

		if !b {
			usersOk = append(usersOk, u)
		}
	}

	// if response body is empty
	if len(usersOk) == 0 {
		fmt.Println("empty body")
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusNoContent)

		return
	}

	// write values getted in to response
	w.Header().Set("content-type", "application/json") // 200
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersOk)

}
