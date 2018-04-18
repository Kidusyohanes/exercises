package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//TODO: implement handlers for your task-related resources here
func (ctx *Context) TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasks, err := ctx.store.GetAll()
		if err != nil {
			http.Error(w, fmt.Sprintf("error getting tasks: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Add(headerContentType, contentTypeJSON)
		if err := json.NewEncoder(w).Encode(tasks); err != nil {
			log.Printf("error encoding JSON: %v", err)
		}
	}
}
