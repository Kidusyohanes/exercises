package main

import (
	"log"
	"os"
	"strings"
)

const usage = `
usage:
	tasks insert "my new task"
	tasks list
	tasks update <task-id> true|false
	tasks purge

`

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

	//TODO: connect to the MySQL server using the information
	//gathered from those environment variables.
	//see https://drstearns.github.io/tutorials/godb/#secconnectingfromagoprogram
	//PRO TIP: the mysql driver in particular has a config
	//struct you can use to build the DSN.
	//see https://godoc.org/github.com/go-sql-driver/mysql#Config
	//and https://godoc.org/github.com/go-sql-driver/mysql#Config.FormatDSN
	//(other drivers may not have something like that)

	//TODO: once connected, create a new tasks.MySQLStore
	//and use it to implement the various commands

	switch command {
	case "insert":
		//TODO: get the new task title from os.Args[2],
		//insert it, and log the new ID or any errors
	case "list":
		//TODO: get all the tasks and log them, one per line,
		//using the following format:
		//<ID>\t<Completed>\t<Title>

		//For example:
		//1  false  get milk
		//2  false  walk the cat

	case "update":
		//TODO: update the task's completed state
		//using os.Args[2] as the task ID
		//and os.Args[3] as the new completed value
		//log the task returned from your store's Update method
		//just like you did in the "list" command
	case "purge":
		//TODO: purge all completed tasks and log
		//how many tasks were deleted
	default:
		logger.Fatal(usage)
	}
}
