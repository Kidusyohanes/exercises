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
  "fmt"
)

type Ctx struct {
  Db models.ZipIndex
}

func (c *Ctx) Handler(w http.ResponseWriter, r *http.Request) {
  /**
    This is here for illustrative purposes only to demonstrate how a server works.
    For the actual exercise, implement the handler inside zip.go in handlers/ folder.
  */
  if r.Method == "GET" {
    query := r.FormValue("q")

    if query == "" {
      http.Error(w, "Query parameter cannot be an empty string.", http.StatusBadRequest)
      return
    }

    fmt.Println(query)

    header := w.Header()
    header.Set("Access-Control-Allow-Origin", "*")
    header.Set("Content-Type", "application/json")

    zList, found := c.Db[query]

    if found {
      jsonStr, _ := json.Marshal(zList)
      w.Write(jsonStr)
    } else {
      w.Write([]byte("[]"))
    }
  } else {
    http.Error(w, "Only GET Method allowed.", http.StatusBadRequest)
  }
}
