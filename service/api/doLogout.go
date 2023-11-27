package api

import (
	"io"
	"strconv"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
If the user is logged,
it will be logout if user not logged method to do nothing.
*/

func (rt *_router) doLogout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var token database.Token

	w.Header().Set("content-type", "text/plain")
	id, _ := strconv.Atoi(ps.ByName("tokenId"))
	token.TokenId = uint64(id)

	for v, user := range UserLogged {
		if user.TokenId == token.TokenId {
			UserLogged[v] = database.Token{}
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "success, user logout action successful\t\n")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "not found, user not logged\t\n")

}
