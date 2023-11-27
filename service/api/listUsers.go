package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"strconv"

	"github.com/julienschmidt/httprouter"
)

/*
this method return the list of users currently registered
*/
func (rt *_router) listUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset := 0

	users, _error := rt.db.GetUsers(r.URL.Query().Get("username"))
	if _error != nil {
		fmt.Println(fmt.Errorf("query error: %w", _error))
		return
	}

	if limit == 0 {
		limit = len(users)
	}
	if offset > len(users) {
		fmt.Println(fmt.Errorf("atoi: %w", _error))
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}
	users = users[offset:min(len(users), limit)]

	if len(users) == 0 {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, "empty body, found nothing")
		return

	} else if _error != nil {
		fmt.Println(fmt.Errorf("atoi: %w", _error))
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	} else {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}

}
