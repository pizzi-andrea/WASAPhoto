package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
given uid then list all user banned by user associated at *uid*
*/
func (rt *_router) listBannedUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var uid_ int
	var err error
	var banned []database.User
	var user *database.User
	var tk *security.Token
	var limit int
	offset := 0

	/*
		Parse URL parameters
	*/
	if uid_, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	if limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil && r.URL.Query().Get("limit") != "" {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	uid := database.Id(uid_)

	/*
		if user id in URL path not exist, then user not found
	*/
	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

	}

	if user == nil {
		w.Header().Set("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	/*

	 */
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	/*
		get banned user and check if not banned
	*/
	if tk.Value != user.Uid {
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	if banned, err = rt.db.GetBanned(uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.WriteHeader(http.StatusInternalServerError)

		return

	}

	if limit == 0 || limit > len(banned) {
		limit = len(banned)
	}

	if offset > len(banned) { //  500 response
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	banned = banned[offset:limit]

	if len(banned) == 0 { //   204 response
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusNoContent)

		return

	} else { //  200 repsonse
		w.Header().Set("content-type", "application/json")
		//   w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(banned); err != nil {
			ctx.Logger.Errorf("%w", err)
			w.Header().Set("content-type", "text/plain") //   500
			w.WriteHeader(ServerError.StatusCode)

		}
	}

}
