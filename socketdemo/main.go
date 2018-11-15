package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type socketStore struct {
	Connections []*websocket.Conn
}

type msg struct {
	Message string `json:"message"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (sockets *socketStore) webSocketConnectionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Incoming..")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// handle the websocket handshake
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to open websocket connection", 401)
		return
	}

	// insert socket connection
	sockets.Connections = append(sockets.Connections, conn)

	// do something with connection
	go echo(conn)
}

func echo(conn *websocket.Conn) {
	for { // infinite loop
		m := msg{
			Message: "Hello",
		}

		err := conn.WriteJSON(m)

		fmt.Printf("%v\n", m)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	mux := http.NewServeMux()

	ctx := socketStore{
		Connections: []*websocket.Conn{},
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		r.FormValue("q")

	})

	mux.HandleFunc("/ws", ctx.webSocketConnectionHandler)
	log.Fatal(http.ListenAndServe(":4001", mux))
}
