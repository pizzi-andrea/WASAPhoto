package api

import (
	"encoding/json"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
)

/*
create new user logged
*/
func (rt *_router) addNewUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "application/json")
	newUser := json.NewDecoder(r.Body)
	var user User
	err := newUser.Decode(&user)
	if err != nil {
		w.WriteHeader(ServerError.StatusCode)

	}
	user.uid = len(Users)
	Users = append(Users, user)
	json.NewEncoder(w).Encode(user)

}
