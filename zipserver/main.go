package main

import (
	"log"
	"net/http"
	"os"

	"github.com/info344-s18/exercises/zipserver/handlers"
	"github.com/info344-s18/exercises/zipserver/models"
)

func main() {
	//TODO: load the zip codes from "zips.csv"
	//build a ZipIndex on the City field
	//and start a web server that responds with
	//all the Zips for a given city name
	f, err := os.Open("zips.csv")
	if err != nil {
		log.Fatalf("error opening zips.csv: %v", err)
	}
	zips, err := models.LoadZips(f, 42613)
	if err != nil {
		log.Fatalf("error parsing CSV: %v", err)
	}
	f.Close()
	log.Printf("loaded %d zips", len(zips))

	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)

	log.Printf("server is listening at http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
