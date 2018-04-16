package tasks

import (
	"database/sql"
	"fmt"
)

const sqlInsertTask = "insert into tasks (title,completed) values (?,?)"
const sqlSelectAll = "select id,title,completed from tasks"
const sqlSelectByID = sqlSelectAll + " where id=?"
const sqlUpdate = "update tasks set completed=? where id=?"
const sqlPurge = "delete from tasks where completed=true"

//MySQLStore represents a tasks.Store backed by MySQL
type MySQLStore struct {
	//db is the open database object
	//this store will use to send queries
	//to the database
	db *sql.DB
}

//NewMySQLStore constructs a new MySQLStore. It will
//panic if the db pointer is nil.
func NewMySQLStore(db *sql.DB) *MySQLStore {
	if db == nil {
		panic("nil database pointer")
	}
	return &MySQLStore{db}
}

//Insert inserts a task into the database, returning
//the inserted Task with its ID field set to the
//new primary key value
func (s *MySQLStore) Insert(task *Task) (*Task, error) {
	//execute the insert task statemen, providing
	//the title and completed values as parameters
	//see https://drstearns.github.io/tutorials/godb/#secguardingagainstsqlinjectionattacks
	results, err := s.db.Exec(sqlInsertTask, task.Title, task.Completed)
	if err != nil {
		return nil, fmt.Errorf("executing insert: %v", err)
	}
	//get the new DBMS-generated primary key value
	id, err := results.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("getting new ID: %v", err)
	}
	//set the ID field of the struct so that callers
	//know what the new ID is
	task.ID = id
	return task, nil
}

//GetAll returns all tasks
func (s *MySQLStore) GetAll() ([]*Task, error) {
	//select all the tasks
	rows, err := s.db.Query(sqlSelectAll)
	if err != nil {
		return nil, fmt.Errorf("selecting: %v", err)
	}
	//ensure the rows get closed, regardless of
	//how this function exits
	defer rows.Close()
	//create a slice of *Task for the results
	var list []*Task
	//while there are rows to read...
	for rows.Next() {
		//create an empty task struct instance
		task := &Task{}
		//scan the row data into the struct fields
		rows.Scan(&task.ID, &task.Title, &task.Completed)
		//append the task to the slice
		list = append(list, task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating over rows: %v", err)
	}
	return list, nil
}

//Get returns a specific task, or ErrNotFound
//if the requested task does not exist
func (s *MySQLStore) Get(id int64) (*Task, error) {
	//since we know we will get only one row back,
	//use QueryRow(). this is a shortcut that eliminates
	//the need to close or iterate over the rows.
	row := s.db.QueryRow(sqlSelectByID, id)
	task := &Task{}
	//with QueryRow(), any errors are surfaced only when you
	//scan the data into the struct fields
	if err := row.Scan(&task.ID, &task.Title, &task.Completed); err != nil {
		//if the error is sql.ErrNoRows, then the
		//request id was not found: return ErrNotFound
		//so that clients can detect that.
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		//otherwise, prefix the error so that callers
		//know what we were trying to do when the error
		//occurred
		return nil, fmt.Errorf("scanning: %v", err)
	}
	return task, nil
}

//Update updates a task, setting only the completed state,
//and returns a copy of the updated Task. It returns
//nil and ErrNotFound if the task ID does not exist.
func (s *MySQLStore) Update(id int64, completed bool) (*Task, error) {
	results, err := s.db.Exec(sqlUpdate, completed, id)
	if err != nil {
		return nil, fmt.Errorf("updating: %v", err)
	}
	affected, err := results.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("getting rows affected: %v", err)
	}
	//if no rows were affected, then the requested
	//ID was not in the database
	if affected == 0 {
		return nil, ErrNotFound
	}
	return s.Get(id)
}

//Purge deletes all completed tasks and returns
//the number of tasks that were deleted
func (s *MySQLStore) Purge() (int64, error) {
	results, err := s.db.Exec(sqlPurge)
	if err != nil {
		return 0, fmt.Errorf("deleting: %v", err)
	}
	return results.RowsAffected()
}
