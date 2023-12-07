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
give user id and put user profile. User profile is object contain
all information on user, in particular:
  - information about user
  - stream photos updated
  - number of photo have been updated
  - number of followers
  - number of following
*/
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var uid_ int
	var err error
	var user *database.User
	var isBan bool
	var follower, following []database.User
	var photos database.StreamPhotos
	var tk *security.Token

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
	/*
		if user id in URL path not exist, then user not found
	*/
	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

	}

	if user == nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Add("content-type", "text/plain") //   404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	/*
		Applay barrear authentication. Only owner user and users not banned by owner can access
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //   401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	/*
		check if user wont access is not beetween banned users from owner
	*/

	if isBan, err = rt.db.IsBanned(user.Uid, tk.Value); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

	}

	if isBan {
		w.Header().Set("content-type", "text/plain") //   403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	/*
		get follower of user
	*/
	if follower, err = rt.db.GetFollowers(uid, "", true); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	/*
		get users following by user
	*/
	if following, err = rt.db.GetFollowing(uid, "", true); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	/*
		get photo stream. Photo stream is composed by photos of other followed users
	*/
	if photos, err = rt.db.GetPhotos(uid, []database.OrderBy{}); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	/*
		put in response body user profile rappresentation
	*/
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK) //   200
	//
	if err = json.NewEncoder(w).Encode(database.Profile{
		User:      *user,
		Stream:    photos,
		Follower:  len(follower),
		Following: len(following),
	}); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //   500
		// w.WriteHeader(ServerError.StatusCode)

	}

}
