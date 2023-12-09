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
give *uid* and *photoId* and get photo associated
*/
func (rt *_router) getComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var uid_ int
	var err error
	var user *database.User
	var tk *security.Token
	var isBan bool
	var photoId_ int
	var photo *database.Photo
	var comment *database.Comment

	//  Parsing URL parameters
	if uid_, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	if photoId_, err = strconv.Atoi(ps.ByName("photoId")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	photoId := database.Id(photoId_)
	uid := database.Id(uid_)

	if !(database.ValidateId(photoId) && database.ValidateId(uid)) {
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(http.StatusBadRequest)

		return

	}

	//  check if path exist
	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	if photo, err = rt.db.GetPhoto(photoId); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	if user == nil || photo == nil {
		w.Header().Add("content-type", "text/plain") //  404
		w.WriteHeader(http.StatusNotFound)

		return

	}

	/*
		Applay barrear authentication. A user can see photos of another user as long as they have not been banned
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //  401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	if isBan, err = rt.db.IsBanned(user.Uid, tk.Value); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	if isBan {
		w.Header().Set("content-type", "text/plain") //  403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}

	if comment, err = rt.db.GetComment(photoId); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(comment)

}
