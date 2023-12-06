package api

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

/*
UploadPhoto method allow users to update their photos in WASAPhoto system. Request MUST formatted like
multipart/form-data. It will parserd values and stored its in db. The photo will be encoding in base64 format
and stored in db.
*/
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var uid int
	var err error
	var user *database.User
	var tk *security.Token
	var file multipart.File

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
		fmt.Println(fmt.Errorf("get exist: %w", err))
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	// if user not exist path is not valid
	if user == nil {
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Not found, user not found")
		return

	}

	/*
		Applay barrear authentication. only owner can post photo in his account
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") // 401
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	/*
		checks if the user who wants post photo is account owner
	*/
	if tk.Value != uint64(uid) {
		w.Header().Set("content-type", "text/plain") // 403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		io.WriteString(w, UnauthorizedToken.Status)
		return
	}
	// parsing body values (image upload operation)
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

	if file, _, err = r.FormFile("imageData"); err != nil {
		fmt.Println(fmt.Errorf("query error: %w", err))
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return

	}

	photo.ImageData, _ = io.ReadAll(file)

	photo.DescriptionImg = r.PostFormValue("descriptionImg")

	if photo.ImageData == nil || len(photo.DescriptionImg) == 0 {
		fmt.Println("error field multipart request empty")
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

	// check if format photo in conform to APIs specifications */
	if !photo.Verify() {
		fmt.Println("Photo not conform to APIs specification")
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, BadRequest.Status)
		return

	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	photo.ImageData = nil // discard img data
	json.NewEncoder(w).Encode(photo)
}
