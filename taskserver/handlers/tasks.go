package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/info344-s18/exercises/tasks/models/tasks"
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
		respond(w, tasks, http.StatusOK)
	case http.MethodPost:
		task := &tasks.Task{}
		if err := json.NewDecoder(r.Body).Decode(task); err != nil {
			http.Error(w, fmt.Sprintf("error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}
		task, err := ctx.store.Insert(task)
		if err != nil {
			http.Error(w, fmt.Sprintf("error inserting: %v", err), http.StatusInternalServerError)
			return
		}
		respond(w, task, http.StatusCreated)
	}
}

func (ctx *Context) SpecificTaskHandler(w http.ResponseWriter, r *http.Request) {
	//gets the last segment of the path
	//reqID := path.Base(r.URL.Path)
}
