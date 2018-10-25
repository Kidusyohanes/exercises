package middleware

import "net/http"

//User represents a user of the system.
type User struct {
	ID       int
	UserName string
}

//GetAuthenticatedUser is a function that returns the
//current user given a request, or nil if the user is
//not currently authenticated. This is just for demo
//purposes: normally you would use your sessions package
//to get the currently authenticated user.
func GetAuthenticatedUser(r *http.Request) (*User, error) {
	return &User{1, "test"}, nil
}

//TODO: define a type for authenticated handler functions
//that take a `*User` as a third parameter

type AuthenticatedHandlerFunc func(w http.ResponseWriter, r *http.Request, u *User)

//TODO: create an adapter function that can adapt an
//authenticated handler function into a regular http
//handler function

func EnsureAuthentication(handlerFunc AuthenticatedHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := GetAuthenticatedUser(r)
		if err != nil {
			http.Error(w, "please sign in", http.StatusUnauthorized)
			return
		}
		handlerFunc(w, r, user)
	}
}

type AuthenticatedMux struct {
	http.ServeMux
}

func NewAuthenticatedMux() *AuthenticatedMux {
	return &AuthenticatedMux{}
}

func (am *AuthenticatedMux) HandleAuthenticatedFunc(pattern string, handlerFunc AuthenticatedHandlerFunc) {
	am.HandleFunc(pattern, EnsureAuthentication(handlerFunc))
}
