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

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var photo_, uid_ int
	var err error
	var msg *database.Comment = &database.Comment{}
	var tk *security.Token
	var user *database.User
	var photo *database.Photo
	var rr bool

	if uid_, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("Atoi::%w\n", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}

	if photo_, err = strconv.Atoi(ps.ByName("photoId")); err != nil {
		ctx.Logger.Errorf("Atoi::%w\n", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)

		return
	}
	// problem ================================================
	photoId := database.Id(photo_)
	uid := database.Id(uid_)
	if err = json.NewDecoder(r.Body).Decode(msg); err != nil {
		ctx.Logger.Errorf("Decode::%w\n", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(BadRequest.StatusCode)
		return

	}
	// =========================================================
	/*
		if !(database.ValidateId(photoId) && database.ValidateId(uid) && msg.Verify()) {
			ctx.Logger.Errorln("Invalid data input")
			w.Header().Set("content-type", "text/plain") //  404
			w.WriteHeader(http.StatusNotFound)
			return

		}
	*/

	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("GetUserFromId::%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if photo, err = rt.db.GetPhoto(photoId); err != nil {
		ctx.Logger.Errorf("GetPhoto::%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if photo == nil || user == nil || msg == nil {
		ctx.Logger.Error("Photo/user not found", err)
		w.Header().Add("content-type", "text/plain") //  404
		w.WriteHeader(http.StatusNotFound)

		return
	}

	/*
		Check if user that wont put follows can do it
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		ctx.Logger.Errorln("Token missing or invalid")
		w.Header().Set("content-type", "text/plain") //  401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	if rr, err = rt.db.IsBanned(user.Uid, tk.Value); err != nil {
		ctx.Logger.Errorf("IsBanned::%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if rr {
		w.Header().Set("content-type", "text/plain") //  403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		return
	}

	if user, err = rt.db.GetUserFromUser(msg.Author.Username); err != nil {
		ctx.Logger.Errorf("Atoi::%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return

	}

	if msg, err = rt.db.PostComment(user.Uid, msg.Text, photo.PhotoId); err != nil {
		ctx.Logger.Errorf("PostComment::%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(http.StatusInternalServerError)

		return

	}

	if msg == nil {
		ctx.Logger.Errorln("Post not created")
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(*msg); err != nil {
		ctx.Logger.Errorf("Encode::%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

	}
	ctx.Logger.Infof("New comment added ")

}
