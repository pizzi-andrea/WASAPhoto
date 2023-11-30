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

/*
If the user does not exist, it will be created,

	and an identifier is returned.
	If the user exists, the user identifier is returned.
*/
type loginUser struct {
	Name database.Username `json:"name"`
}

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	var u loginUser
	var token security.Token

	/*
		Decode values in body request *r
	*/
	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		println(fmt.Errorf("%w", err))
		w.Header().Set("content-type", "text/plain") //400
		w.WriteHeader(BadRequest.Request.Response.StatusCode)
		io.WriteString(w, BadRequest.Status)
		return

	}

	//  check if user is registred
	user, _ := rt.db.GetUsers(u.Name)

	/*
		If user registred, make logged
	*/
	if len(user) == 1 {
		token = security.Token{
			TokenId: user[0].Uid,
			Owner:   user[0].Username,
		}
		w.Header().Set("content-type", "application/json") //201 code
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, "User just exist, log-in action successful")
		security.RecordToken(token)
		json.NewEncoder(w).Encode(token)
		return

	}
	// check db query error
	newUser := database.NewUser(0, u.Name)
	if userQuery, err := rt.db.PostUser(newUser); err != nil {
		w.Header().Set("content-type", "text/plain") // 500 code
		w.WriteHeader(ServerError.StatusCode)
		io.WriteString(w, ServerError.Status)
		return

	} else {
		/*
			if user not registred,  registred it and make logged
		*/

		token = security.Token{
			TokenId: userQuery.GetId(),
			Owner:   userQuery.Username,
		}
		security.RecordToken(token) // 201 code
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, "User create, log-in action successful\n")
		json.NewEncoder(w).Encode(token)
	}

}
