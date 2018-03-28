package handlers

import (
	"fmt"
	"net/http"
)

//TODO: implement an http.Handler that is
//initialized with a ZipIndex, and handles
//HTTP requests where the last segment of the
//resource path is a key into that index.
//respond by JSON-encoding the ZipSlice you get
//from the index given the key

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(headerAccessControlAllowOrigin, "*")
	fmt.Fprintf(w, "Hello, World!")
}
