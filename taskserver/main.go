package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/info344-s18/exercises/tasks/models/tasks"
	"github.com/info344-s18/exercises/taskserver/handlers"
)

func reqEnv(name string) string {
	val := os.Getenv(name)
	if len(val) == 0 {
		log.Fatalf("please set %s variable", name)
	}
	return val
}

func main() {
	//TODO: read the following environment variables.
	//if any are blank, and there is no reasonable default,
	//log a fatal message and exit.
	// - ADDR address web server should listen at (default to ":80")
	// - MYSQL_ADDR address of MySQL server
	// - MYSQL_ROOT_PASSWORD password for MySQL root account
	// - MYSQL_DATABASE name of database we should connect to
	mysqlAddr := reqEnv("MYSQL_ADDR")
	mysqlDB := reqEnv("MYSQL_DATABASE")
	mysqlPwd := reqEnv("MYSQL_ROOT_PASSWORD")
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	//TODO: open a connection to the MySQL database
	//using the values you read above
	config := mysql.Config{
		Addr:   mysqlAddr,
		User:   "root",
		Passwd: mysqlPwd,
		Net:    "tcp",
		DBName: mysqlDB,
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	//TODO: construct a new tasks.Store
	store := tasks.NewMySQLStore(db)

	//TODO: construct a new handlers.Context
	hctx := handlers.NewContext(store)

	//TODO: construct a new mux and add routes
	//for your REST API
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/tasks", hctx.TasksHandler)

	//TODO: start the web server
	log.Printf("server is listening at http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
