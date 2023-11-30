package api

import (
	"encoding/json"
	"io"
	"net/http"

	"strconv"

	"github.com/julienschmidt/httprouter"
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
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var uid int
	var err error
	var user database.User
	var users []database.User
	var follower, following []database.User
	var photos database.StreamPhotos
	var tk *security.Token

	/*
		Parse URL parameters
	*/
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		w.Header().Set("content-type", "plain/text") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	/*
		if user id in URL path not exist, then user not found
	*/
	if user, err = rt.db.GetUser(database.Id(uid)); err != nil {
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "user not found")
		return
	}

	/*
		Applay barrear authentication. Only owner user and users not banned by owner can access
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "plain/text") // 401
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	/*
		get banned user
	*/
	if users, err = rt.db.GetBanned(database.Id(uid)); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	/*
		check if user wont access is not beetween banned users from owner
	*/
	for _, u := range users {
		if tk.Owner == u.Username {
			w.Header().Set("content-type", "plain/text") // 403
			w.WriteHeader(UnauthorizedToken.StatusCode)
			io.WriteString(w, UnauthorizedToken.Status)
			return
		}
	}

	/*
		get follower of user
	*/
	if follower, err = rt.db.GetFollowers(database.Id(uid)); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	/*
		get users following by user
	*/
	if following, err = rt.db.GetFollowing(database.Id(uid)); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	/*
		get photo stream. Photo stream is composed by photos of other followed users
	*/
	if photos, err = rt.db.GetMyStream(database.Id(uid)); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return

	}

	/*
		put in response body user profile rappresentation
	*/
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	io.WriteString(w, "get user profile corresponding to *uid*")
	json.NewEncoder(w).Encode(database.Profile{
		User:      user,
		Stream:    photos,
		Follower:  len(follower),
		Following: len(following),
	})

}
