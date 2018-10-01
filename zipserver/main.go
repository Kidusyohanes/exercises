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
    This is here for illustrative purposes only to demonstrate how a server works.
    For the actual exercise, implement the handler inside zip.go in handlers/ folder.
  */
  if r.Method == "GET" {
    header := w.Header()
    header.Set("Content-Type", "text/html; charset=utf-8")
    w.Write([]byte("Hello"))
  } else {
    http.Error(w, "Only GET allowed.", http.StatusBadRequest)
  }
}
