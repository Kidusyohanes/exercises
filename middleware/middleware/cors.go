package middleware

import (
  "net/http"
)


// TODO:
// Create a middleware that
// adds "Access-Control-Allow-Origin: *" header for every method


// Write the middleware in three flavors

// 1) Wrapping around the whole Mux
// 2) Wrapping around an individual handler function
// 3) As an adapter, returning a handler

// Technique 1
type CorsMW_1 struct {
    MyHandler http.Handler
}

func (c *CorsMW_1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    c.MyHandler.ServeHTTP(w, r)
}

// Technique 2
type CorsMW_2 struct {
    MyHandlerFunc http.HandlerFunc
}

func (c *CorsMW_2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    c.MyHandlerFunc(w, r)
}

// Technique 3
func CorsMW_3(h http.Handler) http.Handler {
  return http.HandlerFunc( func (w http.ResponseWriter, r * http.Request) {
      w.Header().Set("Access-Control-Allow-Origin", "*")
      h.ServeHTTP(w, r)
  })
}
