package main

import (
  "net/http"
  _ "fmt"
  "log"
  _ "github.com/gorilla/websocket"
)

type SocketStore struct{
  // YOUR CODE HERE
}


func WebSocketConnectionHandler(w http.ResponseWriter, r *http.Request) {
  // YOUR CODE HERE

}


func main() {
  mux := http.NewServeMux()
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello"))
  })
	mux.HandleFunc("/ws", WebSocketConnectionHandler)
	log.Fatal(http.ListenAndServe(":4001", mux))
}
