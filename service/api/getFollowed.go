package api

import (
	"encoding/json"
	"net/http"
	"regexp"

	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
give a UID return a list contanings all followers user
*/
func (rt *_router) getFollowed(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var uid_ int
	var err error
	var user *database.User
	var tk *security.Token
	var limit int
	var username database.Username
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
	uid := database.Id(uid_)

	//   validate username format
	username = r.URL.Query().Get("username")
	rr, err := regexp.MatchString("^.*?$", username)
	if !(rr && err == nil && len(username) <= 16) {
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

	/*
		if user id in URL path not exist, then user not found
	*/
	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	if user == nil {
		w.Header().Add("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	/*
		check if token is given and it user is currently logged
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
		ctx.Logger.Errorf("User <%d> not owner. Owner is <%v>\n", tk.Value, user)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	if followed, err := rt.db.GetFollowed(uid, username, true); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.WriteHeader(http.StatusInternalServerError)

		return

	} else {

		if limit == 0 || limit > len(followed) {
			limit = len(followed)
		}

		if offset > len(followed) { //  500 response
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(ServerError.StatusCode)

			return
		}

		followed = followed[offset:limit]

		if len(followed) == 0 { //   204 response
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(http.StatusNoContent)

		} else { //  200 repsonse
			w.Header().Set("content-type", "application/json")
			if err = json.NewEncoder(w).Encode(followed); err != nil {
				ctx.Logger.Errorf("%w", err)
				w.Header().Set("content-type", "text/plain") //   500
				w.WriteHeader(ServerError.StatusCode)

			}
		}

	}

}
