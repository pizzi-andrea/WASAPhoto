package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"pizzi1995517.it/WASAPhoto/service/database"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var uid int
	var err error
	var user *database.User

	// Parsing URL parameters
	if uid, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		fmt.Println(fmt.Errorf("get uid: %w", err))
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

	if user == nil {
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "user not found")
		return

	}

	// parsing body values
	var photo *database.Photo = &database.Photo{
		ImageData: make([]byte, MaxBytePhoto),
	}

	if err = r.ParseMultipartForm(MaxByteFormData); err != nil {
		fmt.Println(fmt.Errorf("multipart: %w", err))
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, BadRequest.Status)
		return
	}

	photo.ImageData = []byte(r.PostFormValue("imageData"))
	photo.DescriptionImg = r.PostFormValue("descriptionImg")

	if photo.ImageData == nil || len(photo.ImageData) == 0 {
		fmt.Println(fmt.Errorf("multipart(photo): %w", err))
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, BadRequest.Status)
		return
	}

	if photo, err = rt.db.PutPhoto(photo.ImageData, photo.DescriptionImg, user.Uid); err != nil {
		fmt.Println(fmt.Errorf("query error: %w", err))
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	if photo == nil {
		fmt.Println(fmt.Errorf("photo error: %w", err))
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	photo.ImageData = nil // discard img data
	json.NewEncoder(w).Encode(photo)
}
