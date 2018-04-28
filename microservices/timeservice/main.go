package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type TimeHandler struct {
	serviceAddr string
}

func (th *TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The time service at %s says the current time is %s",
		th.serviceAddr, time.Now().Format(time.Kitchen))
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	mux := http.NewServeMux()
	mux.Handle("/time/now", &TimeHandler{addr})
	log.Printf("time service instance is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
