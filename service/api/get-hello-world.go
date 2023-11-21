package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/* create new user logged */
func (rt *_router) addNewUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	_, _ = w.Write([]byte("Hello World!"))
}
