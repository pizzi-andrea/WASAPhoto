package api

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	"github.com/julienschmidt/httprouter"
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

var UserLogged []database.Token

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	var u loginUser
	var token database.Token

	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Println(fmt.Errorf("error json parsing %w", err))
		return

	}

	for _, token := range UserLogged {
		if token.Owner == u.Name {
			w.Header().Set("content-type", "text/plain")
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "User just logged")
			return
		}
	}

	users, _ := rt.db.GetUsers(u.Name)
	for _, user := range users {

		if user.Username == u.Name {
			token = database.Token{
				TokenId: user.GetId(),
				Owner:   user.Username,
			}

			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusCreated)
			io.WriteString(w, "User just exist, log-in action successful")
			json.NewEncoder(w).Encode(token)
			UserLogged = append(UserLogged, token)
			return

		}
	}

	newUser := database.NewUser(uint64(len(users)), u.Name)
	if userQuery, err := rt.db.PostUser(newUser); err != nil {
		fmt.Println(fmt.Errorf("error insert values: %w", err))
		return

	} else {

		token = database.Token{
			TokenId: userQuery.GetId(),
			Owner:   userQuery.Username,
		}

		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, "User create, log-in action successful\n")
		json.NewEncoder(w).Encode(token)
		UserLogged = append(UserLogged, token)
	}

}
