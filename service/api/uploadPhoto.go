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
give photo and update it
*/
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var uid_ int
	var err error
	var user *database.User
	var tk *security.Token

	//  Parsing URL parameters
	if uid_, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	uid := database.Id(uid_)
	//  check if path exist
	if user, err = rt.db.GetUserFromId(uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	if user == nil {
		w.Header().Add("content-type", "text/plain") //  404
		w.WriteHeader(http.StatusNotFound)

		return

	}

	/*
		Applay barrear authentication. only owner can post photo in his account
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") //  401
		w.WriteHeader(UnauthorizedError.StatusCode)

		return
	}

	/*
		checks if the user who wants post photo is account owner
	*/
	if tk.Value != uid {
		w.Header().Set("content-type", "text/plain") //  403
		w.WriteHeader(UnauthorizedToken.StatusCode)

		return
	}
	//  parsing body values
	var photo *database.Photo = &database.Photo{
		ImageData: make([]byte, MaxBytePhoto),
	}

	if err = r.ParseMultipartForm(MaxByteFormData); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	photo.ImageData = []byte(r.PostFormValue("imageData"))
	photo.DescriptionImg = r.PostFormValue("descriptionImg")

	if photo.ImageData == nil || len(photo.ImageData) == 0 {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if photo, err = rt.db.PutPhoto(photo.ImageData, photo.DescriptionImg, user.Uid); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	if photo == nil {
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

		return
	}

	/* check if format photo in conform to APIs specifications */
	if !photo.Verify() {
		w.Header().Set("content-type", "text/plain") //  400
		w.WriteHeader(http.StatusBadRequest)

		return

	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	photo.ImageData = nil //  discard img data
	if err = json.NewEncoder(w).Encode(photo); err != nil {
		ctx.Logger.Errorf("%w", err)
		w.Header().Set("content-type", "text/plain") //  500
		w.WriteHeader(ServerError.StatusCode)

	}
}
