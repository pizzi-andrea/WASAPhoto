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
		taken uid of the user who wants to ban and uid of the user to be banned,
	    bans the user(BannedID)
*/
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var result bool
	var from_, to_ int
	var err error
	var tk *security.Token

	/*
		Parse parameters in path
	*/
	if from_, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	if to_, err = strconv.Atoi(ps.ByName("bannedId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	from, to := database.Id(from_), database.Id(to_)

	/*
		Check if path is valid
	*/
	if _, err = rt.db.GetUserFromId(from); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Add("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	if _, err = rt.db.GetUserFromId(to); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Add("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	//   validate format value
	if !(database.ValidateId(from) && database.ValidateId(to)) {
		w.Header().Add("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	/*
		Secure Bearer Authentication, check if user wont put ban is account owner.
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	/*
		Check if user  that want ban user can it. Only user have put ban
		can delete it
	*/
	if tk.Value != from {
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return

	}

	/*
		Try to ban user, if queries fails server response with error message
	*/

	if result, err = rt.db.PutBan(from, to); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)

	}

	if result { //   check now if user follow other user
		var u *database.User
		if u, err = rt.db.GetUserFromId(to); u == nil || err != nil {
			ctx.Logger.Errorf("%w", err)
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(ServerError.StatusCode)

			return
		}
		/*Delete follows  beetween  users */
		if _, err = rt.db.DelFollow(from, to); err != nil {
			ctx.Logger.Errorf("%w", err)
			w.Header().Set("content-type", "text/plain") //   500
			w.WriteHeader(ServerError.StatusCode)

			return

		} //   if user put ban auto. lost the follow
		if _, err = rt.db.DelFollow(to, from); err != nil { //   if user recive ban auto. lost the follow
			ctx.Logger.Errorf("%w", err)
			w.Header().Set("content-type", "text/plain") //   500
			w.WriteHeader(ServerError.StatusCode)

			return
		}
		if _, err = rt.db.DelFollow(from, to); err != nil {
			ctx.Logger.Errorf("%w", err)
			w.Header().Set("content-type", "text/plain") //   500
			w.WriteHeader(ServerError.StatusCode)

			return
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err = json.NewEncoder(w).Encode(*u); err != nil {

			ctx.Logger.Errorf("%w", err)
			w.Header().Set("content-type", "text/plain") //   500
			w.WriteHeader(ServerError.StatusCode)

		}

	} else {
		w.Header().Set("content-type", "text/plain") //  204
		w.WriteHeader(http.StatusNoContent)

	}
}
