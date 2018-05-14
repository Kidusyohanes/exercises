package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//NotificationsHandler handles requests for the /notifications resource
type NotificationsHandler struct {
	notifier *Notifier
}

//NewNotificationsHandler constructs a new NotificationsHandler
func NewNotificationsHandler(notifier *Notifier) *NotificationsHandler {
	return &NotificationsHandler{notifier}
}

//ServeHTTP handles HTTP requests for the NotificationsHandler
func (nh *NotificationsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//NOTE: this is just a simple handler for testing that
	//triggers a new notification anytime this handler
	//receives an HTTP request using any method.
	//In your real server, you will listen for new messages
	//from your MQ server, and pass them to the Notifier as
	//you receive them.
	w.Header().Add("Access-Control-Allow-Origin", "*")
	msg := fmt.Sprintf("Notification pushed from the server at %s", time.Now().Format("15:04:05"))
	nh.notifier.Notify([]byte(msg))
}

//WebSocketsHandler is a handler for WebSocket upgrade requests
type WebSocketsHandler struct {
	notifier *Notifier
	//TODO: add a field for the websocket.Upgrader
	//see https://godoc.org/github.com/gorilla/websocket
	//and https://godoc.org/github.com/gorilla/websocket#Upgrader
}

//NewWebSocketsHandler constructs a new WebSocketsHandler
func NewWebSocketsHandler(notifer *Notifier) *WebSocketsHandler {
	//create, initialize, and return a new WebSocketsHandler
	return nil
}

//ServeHTTP implements the http.Handler interface for the WebSocketsHandler
func (wsh *WebSocketsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("received websocket upgrade request")
	//TODO: upgrade the connection to a WebScoket
	//see https://godoc.org/github.com/gorilla/websocket

	log.Println("adding client to notifier")
	//TODO: add the new WebSocket connection to the Notifier
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		log.Fatal("please set the ADDR environment variable")
	}

	//create a new Notifier
	notifier := NewNotifier()

	//create a mux and add handlers for WebSocket upgrades
	//and triggering notifications
	mux := http.NewServeMux()
	mux.Handle("/websockets", NewWebSocketsHandler(notifier))
	mux.Handle("/notifications", NewNotificationsHandler(notifier))

	log.Printf("server is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
