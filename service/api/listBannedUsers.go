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
given uid then list all user banned by user associated at *uid*
*/
func (rt *_router) listBannedUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var uid int
	var err error
	var banned []database.User
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

	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil && r.URL.Query().Get("limit") != "" {
		w.Header().Set("content-type", "text/plain") // 400
		fmt.Println(fmt.Errorf("query error: %w", err))
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad request, formating error") // 400 response
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
		get banned user and check if not banned
	*/
	if tk.Value != user.Uid {
		w.Header().Set("content-type", "text/plain") // 403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		io.WriteString(w, UnauthorizedToken.Status)
		return
	}

	if banned, err = rt.db.GetBanned(database.Id(uid)); err != nil {
		fmt.Println(fmt.Errorf("internal error: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, ServerError.Status)
		return

	}

	if limit == 0 {
		limit = len(banned)
	}

	if offset > len(banned) { //500 response
		fmt.Println("offser: no offset valid")
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	//banned = banned[offset:min(len(banned), limit)]

	if len(banned) == 0 { // 204 response
		fmt.Println("response: empty")
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, "empty body, found nothing")
		return

	} else { //200 repsonse
		w.Header().Set("content-type", "application/json")
		// w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(banned)
	}

}
