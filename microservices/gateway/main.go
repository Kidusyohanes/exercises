package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const headerUser = "X-User"

//User represents an authenticated user
type User struct {
	ID       int64
	UserName string
}

//GetUser returns the currently-authenticated user,
//or an error if the user is not authenticated. For
//this demo, this function just returns a hard-coded
//test user. In a real gateway, you should use your
//sessions library to get the current session state,
//which contains the currently-authenticated user.
func GetUser(r *http.Request) (*User, error) {
	return &User{
		ID:       1,
		UserName: "TestUser",
	}, nil
}

//RootHandler handles requests for the root resource
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the gateway!")
}

func reqEnv(name string) string {
	val := os.Getenv(name)
	if len(val) == 0 {
		log.Fatalf("please set the %s environment variable", name)
	}
	return val
}

func main() {
	addr := reqEnv("ADDR")

	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)

	log.Printf("server is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
