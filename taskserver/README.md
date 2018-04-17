# Tasks Web API Server Exercise

During our last exercise, you implemented a Store for Tasks backed by a MySQL database, as well as a command-line interface (CLI) program that used that Store. Now it's time to build a web API over that same Tasks Store.

## Design the REST API

Since our data model is very simple, your server will expose an easy-to-use [REST API](https://drstearns.github.io/tutorials/rest/). A REST API exposes a set of **resources** that clients can manipulate. The resource path in the HTTP request names the resource, and the HTTP method indicates what action the client wants to perform on that resource. If the action requires input from the client (e.g., inserting a new resource into a collection), the client provides that input, encoded in JSON, via the HTTP request body.

Design a REST API that supports the following operations. For each operation, list the HTTP resource path, method, and example request body JSON (if any) the client would send to your server to perform that operation. Also list what the successful response status code will be, along with example response body JSON.

- Get all tasks
- Get a specific task by its ID
- Insert a new task
- Update an existing task
- Purge all completed tasks (note that you can define custom HTTP methods)

## Connect to the Database at Startup

Your web server will need to connect to your MySQL database at startup. You can use the same code you wrote in the `main()` function of your CLI program. Use the various `MYSQL_*` environment variables for the connection information, just like you did in the CLI.

After you connect to your database, construct a new `tasks.MySQLStore`, passing the open `*sql.DB` object. You should be able to import that package directly from the `tasks/models/tasks` directory of this repo (no need to copy the files).

## Implement Handler Functions for your API

Implement handler functions for the REST API you designed earlier. Your handler functions will need access to your `tasks.Store` implementation, so use a handler context receiver as described in the [Sharing Values with Go Handlers tutorial](https://drstearns.github.io/tutorials/gohandlerctx/#secreceivers). The handler context field should be of type `tasks.Store` so that your handlers work with _any implementation_ of that interface.

Remember that a handler function will receive all requests for a given resource path _regardless of which HTTP method was used_. If your API allows multiple methods on a given resource, `switch` on `r.Method` to handle the various methods separately. If the client used an unsupported method, respond with an `http.StatusMethodNotAllowed` error status code (the `default:` case).

## Add Handler Functions and Start Your Server

Back in `main.go`, create a mux, add your handler functions, and start your web server listening on whatever address you get from the `ADDR` environment variable (default to `":80"` if that variable is not set). Use `go install` to compile and install your executable.

Ensure that your MySQL server is running, and that you've set the various `MYSQL_*` environment variables. Then start your API server!

## Test it Using Postman

To test your API server interactively, use [Postman](https://www.getpostman.com/). This tool lets you [send HTTP requests](https://www.getpostman.com/docs/v6/postman/sending_api_requests/requests) using any method, resource path, headers, and body. Test your various APIs and ensure that you get back the expected responses.

## Write Automated Tests for Your Handlers

After you do some interactive testing, write [automated unit tests for your handler functions](http://blog.questionable.services/article/testing-http-handlers-go/). Since unit tests focus on just the function you are testing, and not how it integrates with other pieces of your system, you should implement a mock `tasks.Store` that you use in these unit tests. A mock is an implementation of an interface used for automated testing that simply returns hard-coded test values or errors, depending on how it's initialized/configured. 

For example, a mock implementation of `tasks.Store` would look something like this:

```go
type MockStore struct {
    returnErrors bool
}

func (s *MockStore) GetAll() ([]*Task, error) {
    if s.returnErrors {
        return nil, fmt.Errorf("test error")
    }
    return []*Task{ /* a few test tasks*/ }, nil
}
//...other tasks.Store methods...
```

In your automated tests, you can use this `MockStore` as the `tasks.Store` implementation you share with your handler functions. That way your tests don't require a running instance of MySQL, and they focus solely on the behavior of the handler functions. You can also easily trigger error conditions and ensure your handlers respond appropriately.
