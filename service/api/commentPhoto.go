package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

// /users/{uid}/myPhotos/{photoId}/comments/
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var photo_, uid_ int
	var err error
	var msg *database.Comment = &database.Comment{}
	var tk *security.Token
	var user *database.User
	var photo *database.Photo
	var rr bool

	if uid_, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	if photo_, err = strconv.Atoi(ps.ByName("photoId")); err != nil {
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	photoId := database.Id(photo_)
	uid := database.Id(uid_)

	json.NewDecoder(r.Body).Decode(msg)

	if !(database.ValidateId(photoId) && database.ValidateId(uid) && msg.Verify()) {
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return

	}

	if user, err = rt.db.GetUserFromId(uid); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return

	}

	if photo, err = rt.db.GetPhoto(photoId); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return

	}

	if photo == nil || user == nil {
		fmt.Println(fmt.Errorf("not found %w", err))
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Not Found, User not found")
		return
	}

	/*
		Check if user that wont put follows can do it
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") // 401
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	if rr, err = rt.db.IsBanned(user.Uid, tk.Value); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return

	}

	if rr {
		w.Header().Set("content-type", "text/plain") // 403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		io.WriteString(w, UnauthorizedToken.Status)
		return
	}

	if user, err = rt.db.GetUserFromUser(msg.Author.Username); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return

	}

	if msg, err = rt.db.PostComment(user.Uid, msg.Text, photo.PhotoId); err != nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return

	}

	if msg == nil {
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(*msg)

}
