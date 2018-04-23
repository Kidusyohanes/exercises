package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"

	"github.com/info344-s18/exercises/middleware/middleware"
)

const headerContentType = "Content-Type"
const contentTypeJSON = "application/json"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "The current server time is %v",
		time.Now().Format(time.Kitchen))
}

func errorHandler(w http.ResponseWriter, r *http.Request) *middleware.HTTPError {
	if r.Method != "POST" {
		return &middleware.HTTPError{
			Message:    "you must use post",
			StatusCode: http.StatusMethodNotAllowed,
		}
	}
	fmt.Fprintf(w, "Yay you used POST!")
	return nil
}

func usersMeHandler(w http.ResponseWriter, r *http.Request, u *middleware.User) {
	w.Write([]byte(fmt.Sprintf("current user: %d: %s", u.ID, u.UserName)))
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	mux := middleware.NewAuthenticatedMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/time", timeHandler)
	mux.HandleFunc("/error", middleware.HandleHTTPError(errorHandler))
	mux.HandleAuthenticatedFunc("/users/me", usersMeHandler)

	wrappedMux := middleware.NewLogger(middleware.Throttle(mux, redisClient, 3, time.Minute))

	log.Printf("server is listening at http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}
