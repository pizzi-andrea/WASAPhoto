package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
)

/*
create new user logged
*/
func (rt *_router) addNewUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "text/plain")
	_, _ = w.Write([]byte("Hello World!"))
}
