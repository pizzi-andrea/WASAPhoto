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
 */
func (rt *_router) listFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var uid int
	var err error
	var users []database.User
	var tk *security.Token
	var limit int
	offset := 0

	/*
		Parse URL parameters
	*/
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		w.Header().Set("content-type", "plain/text") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil && r.URL.Query().Get("limit") != "" {
		fmt.Println(fmt.Errorf("query error: %w", err))
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad request, formating error") // 400 response
		return
	}

	/*
		if user id in URL path not exist, then user not found
	*/
	if _, err = rt.db.GetUser(database.Id(uid)); err != nil {
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Not found, user not found")
		return
	}

	/*
		Applay barrear authentication. Only owner user and users not banned by owner can access
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "plain/text") // 401
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	/*
		get banned user
	*/
	if users, err = rt.db.GetBanned(database.Id(uid)); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	/*
		check if user wont access is not beetween banned users from owner
	*/
	for _, u := range users {
		if tk.Owner == u.Username {
			w.Header().Set("content-type", "plain/text") // 403
			w.WriteHeader(UnauthorizedToken.StatusCode)
			io.WriteString(w, UnauthorizedToken.Status)
			return
		}
	}

	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil && r.URL.Query().Get("limit") != "" {
		fmt.Println(fmt.Errorf("query error: %w", err))
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	if followers, err := rt.db.GetFollowers(database.Id(uid)); err != nil {
		fmt.Println(fmt.Errorf("internal error: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, ServerError.Status)
		return

	} else {

		if limit == 0 {
			limit = len(followers)
		}

		if offset > len(followers) { //500 response
			fmt.Println("offser: no offset valid")
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(ServerError.StatusCode)
			io.WriteString(w, ServerError.Status)
			return
		}

		followers = followers[offset:min(len(followers), limit)]

		if len(followers) == 0 { // 204 response
			fmt.Println("response: empty")
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(http.StatusNoContent)
			io.WriteString(w, "empty body, found nothing")
			return

		} else { //200 repsonse
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(followers)
		}

	}

}
