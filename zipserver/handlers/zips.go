package handlers

//TODO: implement an http.Handler that is
//initialized with a ZipIndex, and handles
//HTTP requests where the last segment of the
//resource path is a key into that index.
//respond by JSON-encoding the ZipSlice you get
//from the index given the key

import (
  "net/http"
  "exercises/zipserver/models"
  "encoding/json"
)

type Ctx struct {
  Db models.ZipIndex
}

func (c *Ctx) Handler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    query := r.FormValue("q")

    if query == "" {
      http.Error(w, "Query parameter cannot be an empty string.", http.StatusBadRequest)
      return
    }

    header := w.Header()
    header.Set("Content-Type", "application/json")

    zList, found := c.Db[query]

    /*
      Write code here to convert the zList to a JSON and return a reponse

      HINT: use json.Marshal
    */
  } else {
    http.Error(w, "Only GET Method allowed.", http.StatusBadRequest)
  }
}
