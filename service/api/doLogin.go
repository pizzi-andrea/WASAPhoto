package api

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"pizzi1995517.it/WASAPhoto/service/api/security"
	"pizzi1995517.it/WASAPhoto/service/database"
)

type loginUser struct { // struct for json Marshaler
	Name database.Username `json:"name"`
}

/*
If the user does not exist, it will be created,

	and an identifier is returned.
	If the user exists, the user identifier is returned.
*/
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	var u loginUser
	var token security.Token
	var user *database.User

	/*
		Decode values in body request *r
	*/
	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Println(fmt.Errorf("NewDecoder: %w", err))
		w.Header().Set("content-type", "text/plain") //400
		w.WriteHeader(BadRequest.Request.Response.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return

	}

	//  check if user is registred
	if user, err = rt.db.GetUserFromUser(u.Name); err != nil {
		fmt.Println(fmt.Errorf("GetUserFromUser: %w", err))
		w.Header().Set("content-type", "text/plain") // 500 code
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return

	}

	/*
		If user registred, make logged
	*/
	if user != nil {

		token = security.Token{
			Value: user.Uid,
		}

		//io.WriteString(w, "User just exist, log-in action successful")
		security.RecordToken(token)
		json.NewEncoder(w).Encode(token)
		w.Header().Set("content-type", "application/json") //201 code
		//w.WriteHeader(http.StatusCreated)
		return

	}
	// check db query error

	if user, err = rt.db.PostUser(u.Name); err != nil || user == nil {

		fmt.Println(fmt.Errorf("PostUser:%w", err))
		w.Header().Set("content-type", "text/plain") // 500 code
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return

	}

	/*
		if user not registred,  registred it and make logged
	*/

	token = security.Token{
		Value: user.Uid,
	}

	security.RecordToken(token) // 201 code
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)

}
