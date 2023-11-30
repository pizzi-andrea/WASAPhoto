/*
package that implement simple scheme of authentication based on barear authentication
*/
package security

import (
	"net/http"
	"strconv"
	"strings"
)

var tokenRegistred map[string]int = make(map[string]int) // list token registred in the system. The token is registred in login action and deleted in logout action

func RecordToken(token Token) (status bool) {
	if token.Owner == "" {
		return false
	}

	if _, exitst := tokenRegistred[token.Owner]; !exitst {
		tokenRegistred[token.Owner] = int(token.TokenId)
	}

	return true

}

func GetTokensRec() (t []Token) {

	for k, v := range tokenRegistred {
		t = append(t, Token{
			TokenId: uint64(v),
			Owner:   k,
		})

	}
	return
}

func GetToken(o string) *Token {
	if v, ok := tokenRegistred[o]; ok {
		return &Token{
			TokenId: uint64(v),
			Owner:   o,
		}

	}
	return nil
}

func TokenIn(token Token) (reponse bool) {
	if v, ok := tokenRegistred[token.Owner]; ok {
		return v == int(token.TokenId)
	}
	return false
}

/*
Simple implementations of barrear authentications methods.
Given in input http request and exstract token if provided and return it. Oterwhise
return nil.
*/
func BarrearAuth(r *http.Request) (tokenValid *Token) {
	if r == nil {
		return nil
	}
	var id int
	var err error
	var owner string

	s := strings.Split(r.Header.Get("Authorization"), " ")
	if len(s) != 2 {
		return nil
	}
	typeAuth := s[0]
	cred := strings.Split(s[1], ";")

	if len(cred) != 2 {
		return nil
	}

	if id, err = strconv.Atoi(cred[0]); err != nil {
		return nil
	}
	owner = cred[1]

	if owner == "" || typeAuth == "" || typeAuth != "Bearer" {
		return nil
	}

	return &Token{
		uint64(id),
		owner,
	}

}
