package api

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*
testing method on root path
*/
func (rt *_router) hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "<h1>WASAPhoto Version 0.1</h1>\n\n <i>developed with ❤️ by: pizzi.1995517@studenti.uniroma1.it</i>")
	io.WriteString(w, "<p> Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto!<br><br> WASAPhoto is web application based on RESTfull architecture to allows upload your photos directly from your PC, and they will be visible to everyone following you.")

}
