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
	follower := 0
	following := 0
	users := []database.User{}
	var user database.User
	var photos database.StreamPhotos

	uid, _error := strconv.Atoi(ps.ByName("uid"))

	if _error != nil {
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
	}

	for _, user = range users {
		if user.GetId() == uint64(uid) {

			// get stream of photos
			w.Header().Add("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "get user profile corresponding to *uid*")
			json.NewEncoder(w).Encode(database.Profile{
				User:      user,
				Stream:    photos,
				Follower:  follower,
				Following: following,
			})
			return
		}
	}
	w.Header().Add("content-type", "text/plain")
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "user not found")

}
