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

/*
give *uid* and *photoId* and get photo associated
*/
func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var uid_ int
	var err error
	var user *database.User
	var tk *security.Token
	var isBan bool
	var photoId_ int
	var photo *database.Photo

	// Parsing URL parameters

	if uid_, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		fmt.Println(fmt.Errorf("get uid: %w", err))
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, BadRequest.Status)
		return
	}
	if photoId_, err = strconv.Atoi(ps.ByName("photoId")); err != nil {
		fmt.Println(fmt.Errorf("get uid: %w", err))
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, BadRequest.Status)
		return
	}

	photoId := database.Id(photoId_)
	uid := database.Id(uid_)

	// check if path exist
	if user, err = rt.db.GetUserFromId(uid); err != nil {
		fmt.Println(fmt.Errorf("user exist: %w", err))
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return
	}

	if photo, err = rt.db.GetPhoto(photoId); err != nil {
		fmt.Println(fmt.Errorf("error GetPhoto: %w", err))
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
		Applay barrear authentication. A user can see photos of another user as long as they have not been banned
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		w.Header().Set("content-type", "text/plain") // 401
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	if isBan, err = rt.db.IsBanned(user.Uid, tk.Value); err != nil {
		fmt.Println(fmt.Errorf(" %w", err))
		w.Header().Set("content-type", "text/plain") // 500
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status+"TOKKOT")
		return
	}

	if isBan {
		w.Header().Set("content-type", "text/plain") // 403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		io.WriteString(w, UnauthorizedToken.Status)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(*photo)
}
