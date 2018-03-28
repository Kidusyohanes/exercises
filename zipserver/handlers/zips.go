package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/info344-s18/exercises/zipserver/models"
)

//ZipIndexHandler represents a zip index handler
type ZipIndexHandler struct {
	index models.ZipIndex
}

func NewZipIndexHandler(index models.ZipIndex) *ZipIndexHandler {
	return &ZipIndexHandler{
		index: index,
	}
}

func (zih *ZipIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// /zips/city/city-name
	key := path.Base(r.URL.Path)
	key = strings.ToLower(key)

	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add(headerAccessControlAllowOrigin, "*")
	zips := zih.index[key]
	json.NewEncoder(w).Encode(zips)
}

//TODO: implement an http.Handler that is
//initialized with a ZipIndex, and handles
//HTTP requests where the last segment of the
//resource path is a key into that index.
//respond by JSON-encoding the ZipSlice you get
//from the index given the key

//RootHandler handles requests for the root resource
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(headerAccessControlAllowOrigin, "*")
	fmt.Fprintf(w, "Try requesting /zips/city/seattle")
}
