package api

import (
	"encoding/json"
	"io"
	"net/http"

	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
this method return the list of users currently registered
*/
func (rt *_router) listUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var users, response []database.User // missing-db

	limit, _error := strconv.Atoi(ps.ByName("limit"))

	if _error != nil {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	username := ps.ByName("username")
	w.Header().Set("content-type", "application/json")

	for _, user := range users[0:min(len(users), limit)] {
		if user.Username == username || username != "" {
			response = append(response, user)
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
