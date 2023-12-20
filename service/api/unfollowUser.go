package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/reqcontext"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
gived uid and *followedId* then remove follower *followerId* from user followers
*/
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var uid_ int
	var err error
	var uidUnfoll_ int
	var isFollower bool
	var tk *security.Token

	/*
		Parse URL parameters
	*/
	if uid_, err = strconv.Atoi(ps.ByName("followerId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	if uidUnfoll_, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	/*
		if folow not exist path URL path not exist
	*/
	uid := database.Id(uid_)
	uidUnfoll := database.Id(uidUnfoll_)

	if isFollower, err = rt.db.IsFollower(uid, uidUnfoll); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500 code
		w.WriteHeader(ServerError.StatusCode)
		return

	}

	if !isFollower {
		w.Header().Add("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return

	}

	/*
		Applay barrear authentication. only owner can deleted follow
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)
		return
	}

	/*
		checks if the user who wants make unfollow is owner
	*/

	if tk.Value != uid {
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	if _, err = rt.db.DelFollow(uid, uidUnfoll); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500 code
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	w.Header().Set("content-type", "text/plain") //   204
	w.WriteHeader(http.StatusNoContent)

}
