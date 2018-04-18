package handlers

import "github.com/info344-s18/exercises/tasks/models/tasks"

type Context struct {
	store tasks.Store
}

func NewContext(store tasks.Store) *Context {
	if store == nil {
		panic("nil store pointer")
	}
	return &Context{store}
}
