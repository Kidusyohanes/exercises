package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/info344-s18/prep/tasks/models/tasks"
)

const usage = `
usage:
	tasks insert "my new task"
	tasks list
	tasks update <task-id> true|false
	tasks purge

`

func reqEnv(name string) string {
	val := os.Getenv(name)
	if len(val) == 0 {
		log.Fatalf("please set the %s environment variable", name)
	}
	return val
}

func insert(store tasks.Store, logger *log.Logger) {
	if len(os.Args) < 3 {
		logger.Fatal(usage)
	}
	task := tasks.NewTask(os.Args[2])
	task, err := store.Insert(task)
	if err != nil {
		logger.Fatalf("error inserting task: %v", err)
	}
	logger.Println(task.ID)
}

func list(store tasks.Store, logger *log.Logger) {
	list, err := store.GetAll()
	if err != nil {
		logger.Fatalf("error getting tasks: %v", err)
	}
	for _, task := range list {
		printTask(task, logger)
	}
}

func update(store tasks.Store, logger *log.Logger) {
	if len(os.Args) < 4 {
		logger.Fatalf(usage)
	}
	taskID, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		logger.Fatalf("`%s` is not a valid task ID", os.Args[2])
	}
	completed, err := strconv.ParseBool(os.Args[3])
	if err != nil {
		logger.Fatalf("`%s` cannot be interpreted as true or false", os.Args[3])
	}
	task, err := store.Update(taskID, completed)
	if err != nil {
		logger.Fatalf("error updating: %v", err)
	}
	printTask(task, logger)
}

func purge(store tasks.Store, logger *log.Logger) {
	numDeleted, err := store.Purge()
	if err != nil {
		logger.Fatalf("error purging: %v", err)
	}
	logger.Printf("%d deleted", numDeleted)
}

func printTask(task *tasks.Task, logger *log.Logger) {
	logger.Printf("%d\t%v\t%s", task.ID, task.Completed, task.Title)
}

func main() {
	//create a new logger with no date/time prefix.
	//Use this to write responses back to the terminal.
	//Use logger.Fatalf() to log a fatal message and
	//exit with a non-zero status code.
	//use logger.Printf() to write other messages.
	logger := log.New(os.Stdout, "", 0)
	if len(os.Args) < 2 {
		logger.Fatal(usage)
	}

	//command will be one of:
	//insert | list | update | purge
	command := strings.ToLower(os.Args[1])

	//TODO: read the following required env vars.
	//if they are not defined, use logger.Fatalf()
	//to tell the user to define them and exit
	// - MYSQL_ADDR = network address of the MySQL server (127.0.0.1:3306)
	// - MYSQL_ROOT_PASSWORD = password for the root user account
	// - MYSQL_DATABASE = name of database containing our tasks table
	dbAddr := reqEnv("MYSQL_ADDR")
	dbPwd := reqEnv("MYSQL_ROOT_PASSWORD")
	dbName := reqEnv("MYSQL_DATABASE")

	//TODO: connect to the MySQL server using the information
	//gathered from those environment variables.
	//see https://drstearns.github.io/tutorials/godb/#secconnectingfromagoprogram
	//PRO TIP: the mysql driver in particular has a config
	//struct you can use to build the DSN.
	//see https://godoc.org/github.com/go-sql-driver/mysql#Config
	//and https://godoc.org/github.com/go-sql-driver/mysql#Config.FormatDSN
	//(other drivers may not have something like that)
	config := mysql.Config{
		Addr:   dbAddr,
		User:   "root",
		Passwd: dbPwd,
		DBName: dbName,
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	//TODO: once connected, create a new tasks.MySQLStore
	//and use it to implement the various commands
	store := tasks.NewMySQLStore(db)

	switch command {
	case "insert":
		insert(store, logger)
	case "list":
		list(store, logger)
	case "update":
		update(store, logger)
	case "purge":
		purge(store, logger)
	default:
		logger.Fatal(usage)
	}
}
