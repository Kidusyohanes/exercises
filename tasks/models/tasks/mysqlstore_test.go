package tasks

import (
	"fmt"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAll(t *testing.T) {
	//create a new sql mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sql mock: %v", err)
	}
	//ensure it's closed at the end of the test
	defer db.Close()

	//construct a new task and a set of rows to return
	expectedTask := &Task{
		ID:        1,
		Title:     "test task",
		Completed: true,
	}
	rows := sqlmock.NewRows([]string{"id", "title", "completed"})
	rows.AddRow(expectedTask.ID, expectedTask.Title, expectedTask.Completed)

	//tell sqlmock that we expect the function to execute a
	//a particular SQL query, and that it should return the
	//rows we constructed above
	mock.ExpectQuery(sqlSelectAll).WillReturnRows(rows)

	//construct a new MySQLStore using the mock db
	store := NewMySQLStore(db)

	//call the GetAll() method
	tasks, err := store.GetAll()
	//we shouldn't get an error
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	//but we should get back our one task
	if len(tasks) != 1 {
		t.Errorf("incorrect number of tasks returned: expected %d but got %d",
			1, len(tasks))
	}
	if tasks[0].Title != expectedTask.Title {
		t.Errorf("incorrect task title: expected %s but got %s",
			expectedTask.Title, tasks[0].Title)
	}

	//tell the mock to return an error this time
	mock.ExpectQuery(sqlSelectAll).WillReturnError(fmt.Errorf("test DMBS error"))
	tasks, err = store.GetAll()
	if err == nil {
		t.Errorf("did not receive expected error")
	}

	//ensure we didn't have any unmet expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet sqlmock expectations: %v", err)
	}
}
