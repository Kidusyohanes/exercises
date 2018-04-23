package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const headerContentType = "Content-Type"
const contentTypeJSON = "application/json"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The current server time is %v",
		time.Now().Format(time.Kitchen))
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/time", timeHandler)

	log.Printf("server is listening at http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
