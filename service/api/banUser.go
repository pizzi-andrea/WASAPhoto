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
		taken uid of the user who wants to ban and uid of the user to be banned,
	    bans the user(BannedID)
*/
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var result bool
	var from, to int
	var err error
	var tk *security.Token

	/*
		Parse parameters in path
	*/
	if from, err = strconv.Atoi(ps.ByName("uid")); err != nil {
		fmt.Println(fmt.Errorf("atoi:%w", err))
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	if to, err = strconv.Atoi(ps.ByName("bannedId")); err != nil {
		fmt.Println(fmt.Errorf("atoi:%w", err))
		w.Header().Set("content-type", "text/plain") // 400
		w.WriteHeader(BadRequest.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return
	}

	/*
		Check if path is valid
	*/
	if _, err = rt.db.GetUserFromId(database.Id(from)); err != nil {
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Not found, user not found")
		return
	}

	if _, err = rt.db.GetUserFromId(database.Id(to)); err != nil {
		fmt.Println(fmt.Errorf("GetUserFromId:%w", err))
		w.Header().Add("content-type", "text/plain") // 404
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Not found, user not found")
		return
	}

	/*
		Secure Bearer Authentication, check if user wont put ban is account owner.
	*/
	if tk = security.BarrearAuth(r); tk == nil || !security.TokenIn(*tk) {
		fmt.Println(fmt.Errorf("BarrearAuth/TokenIn:%w", err))
		w.Header().Set("content-type", "text/plain") // 401
		w.WriteHeader(UnauthorizedError.StatusCode)
		io.WriteString(w, UnauthorizedError.Status)
		return
	}

	/*
		Check if user  that want ban user can it. Only user have put ban
		can delete it
	*/
	if tk.Value != database.Id(from) {
		fmt.Println(fmt.Errorf("BarrearAuth/TokenIn:%w", err))
		w.Header().Set("content-type", "text/plain") // 403
		w.WriteHeader(UnauthorizedToken.StatusCode)
		io.WriteString(w, UnauthorizedToken.Status)
		return

	}

	/*
		Try to ban user, if queries fails server response with error message
	*/

	if result, err = rt.db.PutBan(database.Id(from), database.Id(to)); err != nil {
		fmt.Println(fmt.Errorf("PutBan:%w", err))
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
	}

	if result { // check now if user follow other user
		var u *database.User
		if u, err = rt.db.GetUserFromId(database.Id(to)); u == nil || err != nil {
			fmt.Println(fmt.Errorf("GetUserFromId:%w", err))
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(ServerError.StatusCode)
			io.WriteString(w, ServerError.Status)
			return
		}
		/*Delete follows  beetween  users */
		if _, err = rt.db.DelFollow(database.Id(from), database.Id(to)); err != nil {
			fmt.Println(fmt.Errorf("DelFollow: %w", err))
			w.Header().Set("content-type", "text/plain") // 500
			w.WriteHeader(ServerError.StatusCode)
			io.WriteString(w, ServerError.Status)
			return

		} // if user put ban auto. lost the follow
		rt.db.DelFollow(database.Id(to), database.Id(from)) // if user recive ban auto. lost the follow

		if _, err = rt.db.DelFollow(database.Id(from), database.Id(to)); err != nil {
			fmt.Println(fmt.Errorf("DelFollow: %w", err))
			w.Header().Set("content-type", "text/plain") // 500
			w.WriteHeader(ServerError.StatusCode)
			io.WriteString(w, ServerError.Status)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(*u)

	} else {
		w.Header().Set("content-type", "text/plain") //204
		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, "empty response, user just banned")
	}
}
