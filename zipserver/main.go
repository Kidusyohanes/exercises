package main

/*
  Import relevant packages.
  HINT: You may need encoding/json package.
*/
import (
  "net/http"
)

func main() {
	//TODO: load the zip codes from "zips.csv"
	//build a ZipIndex on the City field
	//and start a web server that responds with
	//all the Zips for a given city name

  //create a new mux (router)
  //the mux calls different functions for
  //different resource paths
  mux := http.NewServeMux()

  // register your handler function
  mux.HandleFunc("/", handleFunc)

  http.ListenAndServe(":4000", mux)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
  /**
    Modify this handler function to
    1) return a JSON of Zipcode objects
    2) server CORS requests
  */
  if r.Method == "GET" {
    header := w.Header()
    header.Set("Content-Type", "text/html; charset=utf-8")
    w.Write([]byte("Hello"))
  } else {
    http.Error(w, "Only GET allowed.", http.StatusBadRequest)
  }
}
