package api

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// liveness is an HTTP handler that checks the API server status. If the server cannot serve requests (e.g., some
// resources are not ready), this should reply with HTTP Status 500. Otherwise, with HTTP Status 200
func (rt *_router) liveness(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	/* Example of liveness check:*/
	if err := rt.db.Ping(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "<h1 style='color:red;'> All system components are active with success </h1> ")
		return
	}

	w.Header().Set("content-type", "text/html")
	io.WriteString(w, "<h1 style='color:green;'> All system components are active with success </h1> ")
}
