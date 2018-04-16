package tasks

import "errors"

//ErrNotFound is returned from methods like Get()
//and Update() when the specific task ID is not found
var ErrNotFound = errors.New("record not found")

//Store represents a store for Tasks
type Store interface {
	//Insert inserts a task into the database, returning
	//the inserted Task with its ID field set to the
	//new primary key value
	Insert(task *Task) (*Task, error)

	//GetAll returns all tasks
	GetAll() ([]*Task, error)

	//Get returns a specific task
	Get(id int64) (*Task, error)

	//Update updates a task, setting only the completed state,
	//and returns a copy of the updated Task
	Update(id int64, completed bool) (*Task, error)

	//Purge deletes all completed tasks and
	//returns the number that were deleted
	Purge() (int64, error)
}
