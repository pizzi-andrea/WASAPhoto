package api

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
)

// getContextReply is an example of HTTP endpoint that returns "Hello World!" as a plain text. The signature of this
// handler accepts a reqcontext.RequestContext (see httpRouterHandler).
func (rt *_router) getContextReply(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "<h1>WASAPhoto Version 0.1</h1>\n\n <i>developed with ❤️ by: pizzi.1995517@studenti.uniroma1.it</i>")
	io.WriteString(w, "<p> Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto!<br><br> WASAPhoto is web application based on RESTfull architecture to allows upload your photos directly from your PC, and they will be visible to everyone following you.")
	ctx.Logger.Info(" Hello world by WASAPhoto!")

}
