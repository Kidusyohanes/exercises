package tasks

//Task represents a task in the database
type Task struct {
	ID        int64
	Title     string
	Completed bool
}
