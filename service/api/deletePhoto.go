package api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

// deletePhoto permit to photo owner to delete updated photo
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var uid int
	var photoId int
	var err error
	var user *database.User
	var photo *database.Photo
	var tk *security.Token

	// Parsing URL parameters in path
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		fmt.Println(fmt.Errorf("get uid: %w", err))
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, BadRequest.Status)
		return
	}

	if photoId, err = strconv.Atoi(ps.ByName("photoId")); err != nil {
		fmt.Println(fmt.Errorf("get photoId: %w", err))
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, BadRequest.Status)
		return
	}
	// check if path exist
	if user, err = rt.db.GetUserFromId(database.Id(uid)); err != nil {
		fmt.Println(fmt.Errorf("user exist: %w", err))
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}
	if photo, err = rt.db.GetPhoto(database.Id(photoId)); err != nil {
		fmt.Println(fmt.Errorf("user exist: %w", err))
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	if user == nil || photo == nil {
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "user not found")
		return

	}

	/*
		Applay barrear authentication.
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") // 401
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	/*
		checks if the user who wants delete photo is the owner
	*/
	if tk.Value != uint64(uid) {
		w.Header().Set("content-type", "text/plain") // 403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		io.WriteString(w, UnauthorizedToken.Status)
		return
	}

	// delete photo
	if _, err = rt.db.DelPhoto(photo.PhotoId); err != nil {
		fmt.Println(fmt.Errorf("user exist: %w", err))
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return

	}

	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusNoContent)

}
