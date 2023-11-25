package api

import (
	"encoding/json"
	"io"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
If the user does not exist, it will be created,

	and an identifier is returned.
	If the user exists, the user identifier is returned.
*/
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var users []database.User //

	var str_username string
	username, _error := r.GetBody()
	if _error != nil {
		//TODO
		return
	}

	json.NewDecoder(username).Decode(&str_username)
	for _, user := range users {
		if user.Username == str_username {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusCreated)
			io.WriteString(w, "User log-in action successful")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	newUser := database.NewUser(uint64(len(users)), str_username)

	// serialize new user

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "User log-in action successful")
	json.NewEncoder(w).Encode(newUser)

}
