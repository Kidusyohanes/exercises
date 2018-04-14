package tasks

import (
	"database/sql"
)

//MySQLStore represents a tasks.Store backed by MySQL
type MySQLStore struct {
	//TODO: add field(s) you need
}

//NewMySQLStore constructs a new MySQLStore. It will
//panic if the db pointer is nil.
func NewMySQLStore(db *sql.DB) *MySQLStore {
	panic("TODO: implement this function")
}

//Insert inserts a task into the database, returning
//the inserted Task with its ID field set to the
//new primary key value
func (s *MySQLStore) Insert(task *Task) (*Task, error) {
	panic("TODO: implement this function")
}

//GetAll returns all tasks
func (s *MySQLStore) GetAll() ([]*Task, error) {
	panic("TODO: implement this function")
}

//Get returns a specific task
func (s *MySQLStore) Get(id int64) (*Task, error) {
	panic("TODO: implement this function")
}

//Update updates a task, setting only the completed state,
//and returns a copy of the updated Task
func (s *MySQLStore) Update(id int64, completed bool) (*Task, error) {
	panic("TODO: implement this function")
}

//Purge deletes all completed tasks and returns
//the number of tasks that were deleted
func (s *MySQLStore) Purge() (int64, error) {
	panic("TODO: implement this function")
}
