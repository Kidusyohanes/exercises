package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"
	"path"
)

const headerCORS = "Access-Control-Allow-Origin"

func identiconHandler(w http.ResponseWriter, r *http.Request) {
	//GET /identicon/dave
	name := path.Base(r.URL.Path)
	if len(name) == 0 {
		name = "World"
	}
	img := identicon(name)
	w.Header().Add("Content-Type", "image/png")
	w.Header().Add(headerCORS, "*")
	png.Encode(w, img)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		name = "World"
	}
	w.Header().Add(headerCORS, "*")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/identicon/", identiconHandler)

	log.Printf("server is listening at http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
