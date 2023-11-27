package api

import (
	"encoding/json"
	"io"
	"net/http"

	"strconv"

	"github.com/julienschmidt/httprouter"
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

	uid, _error := strconv.Atoi(ps.ByName("uid"))
	if _error != nil {
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	if user, err := rt.db.GetUser(database.Id(uid)); err != nil {
		w.Header().Add("content-type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "user not found")
		return
	} else {
		follower, _error := rt.db.GetFollower(database.Id(uid))
		following, _error := rt.db.GetFollowing(database.Id(uid))
		photos, _error := rt.db.GetMyStream(database.Id(uid))

		if _error != nil {
			w.WriteHeader(ServerError.StatusCode)
			io.WriteString(w, ServerError.Status)
		}

		// get stream of photos
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "get user profile corresponding to *uid*")
		json.NewEncoder(w).Encode(database.Profile{
			User:      user,
			Stream:    photos,
			Follower:  len(follower),
			Following: len(following),
		})
	}

}
