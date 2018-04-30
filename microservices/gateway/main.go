package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"sync"
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
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) == 0 {
		return nil, fmt.Errorf("not authenticated")
	}
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

//NewServiceProxy returns a new ReverseProxy
//for a microservice given a comma-delimited
//list of network addresses
func NewServiceProxy(addrs string) *httputil.ReverseProxy {
	splitAddrs := strings.Split(addrs, ",")
	nextAddr := 0
	mx := sync.Mutex{}

	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = "http"
			mx.Lock()
			r.URL.Host = splitAddrs[nextAddr]
			nextAddr = (nextAddr + 1) % len(splitAddrs)
			mx.Unlock()

			r.Header.Del(headerUser)
			user, err := GetUser(r)
			if err != nil {
				return
			}
			userJSON, _ := json.Marshal(user)
			r.Header.Set(headerUser, string(userJSON))
		},
	}
}

func main() {
	addr := reqEnv("ADDR")
	timeServiceAddrs := reqEnv("TIME_ADDRS")
	helloServiceAddrs := reqEnv("HELLO_ADDRS")

	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)
	mux.Handle("/time/now", NewServiceProxy(timeServiceAddrs))
	mux.Handle("/hello", NewServiceProxy(helloServiceAddrs))

	log.Printf("server is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
