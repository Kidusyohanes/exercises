# Tasks CLI Exercise

This directory contains the starter files for a `tasks` command-line program that you will implement. This program will allow users to insert, select, update, and delete tasks stored in a relational database management system (RDBMS) such as [MySQL](https://hub.docker.com/_/mysql/).

By implementing this, you will learn how to interact with an RDBMS from Go. This is the first step towards building an HTTP API server that interacts with an RDBMS.

## Define Your Schema

As described in the [Bootstrapping the Database section of the tutorial](https://drstearns.github.io/tutorials/godb/#secbootstrappingthedatabaseschema), we can create a new Docker container image that extends the base MySQL image to automatically create our schema when the container instance starts. Any `.sql` file copied into the `/docker-entrypoint-initdb.d/` directory of your container image will be executed against the MySQL Server when the container runs.

To begin, open the `db/schema.sql` file and add SQL statement(s) to [create a table](https://dev.mysql.com/doc/refman/5.7/en/create-table.html) that can store your tasks. Refer to the fields in the `Task` struct in `models/tasks/task.go` to see what data you need to store.

## Build Your Container Image

Then implement the `db/Dockerfile` as described in [the tutorial](https://drstearns.github.io/tutorials/godb/#secbootstrappingthedatabaseschema), and build your new Docker container image. You might want to write a quick bash script to do this in case you need to rebuild after fixing something in your `schema.sql` file.

## Run an Instance

The MySQL Docker container image uses a few environment variables to setup the root user's password and the default database. Set those now:

```bash
# random password for the root user
export MYSQL_ROOT_PASSWORD=$(openssl rand -base64 18)

# default database name
export MYSQL_DATABASE=tasksex
```

Now run a detached instance of your container image, publishing port `127.0.0.1:3306` on the host to port `3306` in the container, and forwarding these two environment variables:

```bash
docker run -d \
-p 127.0.0.1:3306:3306 \
--name tasksdb \
-e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
-e MYSQL_DATABASE=$MYSQL_DATABASE \
your-dh-name/your-container-name
```

You might want to put that into a bash script so that you can re-run it later without having to re-type all of that!

Use `docker ps -a` to ensure that your container is still running and didn't exit right away. If it did exit, use `docker logs` to see the error messages. If it exited shortly after running, you probably have a syntax error in your `db/schema.sql` file.

## Implement MySQLStore

The `models/tasks/task.go` file defines what a single task looks like. The `models/tasks/store.go` file defines an interface for a store of these tasks. The `models/tasks/mysqlstore.go` contains the stubs for an implementation of this store interface backed by your MySQL database.

Implement the various functions in `models/tasks/mysqlstore.go`. Note that the `*sql.DB` parameter passed to `NewMySQLStore()` is your connection to the database. You can use that to [execute insert/update/delete statements](https://golang.org/pkg/database/sql/#DB.Exec), or [select rows](https://golang.org/pkg/database/sql/#DB.Query).

Refer to the tutorial for examples of [inserting new rows](https://drstearns.github.io/tutorials/godb/#secinsertingandgettingautoassignedids) and [selecting rows](https://drstearns.github.io/tutorials/godb/#secselectingrows).

## Connect to the MySQL Server in Main

The `main.go` file contains the `main()` entry point function. There is some code in there already, but you need to implement the rest according to the `TODO:` comments. Note that the `logger` variable created at the top can be used to write fatal messages and exit with a non-zero status code (`log.Fatalf()`), or write other messages without exiting (`log.Printf()`).

