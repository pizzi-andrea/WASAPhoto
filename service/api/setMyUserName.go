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
assign username or update old username with new username.
The username to set is in the body request
*/

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var users []database.User //missing-db
	w.Header().Set("content-type", "text/plain")

	uid, _error := strconv.Atoi(ps.ByName("uid"))

	newUsername := json.NewDecoder(r.Body)

	var username database.Username
	newUsername.Decode(&username)

	if _error != nil || username == "" {
		return
	}

	for _, user := range users {
		if user.GetId() == uint64(uid) {
			user.Username = username
			w.WriteHeader(http.StatusNoContent)
			io.WriteString(w, "success, assigned")
			return

		}
	}

	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "user not found")

}
