package handlers;

import com.google.gson.Gson;
import io.undertow.server.HttpHandler;
import io.undertow.server.HttpServerExchange;
import io.undertow.util.Headers;
import io.undertow.util.StatusCodes;
import models.SQLStore;
import models.Task;

import java.io.InputStreamReader;
import java.io.Reader;
import java.sql.SQLException;

public class TasksHandler implements HttpHandler {

    // Reference to the SQLStore that we will be using to do any
    // interactions with the database.
    private SQLStore store;

    public TasksHandler(SQLStore store) throws IllegalArgumentException {
        if (store == null) {
            throw new IllegalArgumentException("sql store cannot be null");
        }
        this.store = store;
    }

    public void handleRequest(HttpServerExchange exchange) {
        String method = exchange.getRequestMethod().toString();
        if (method.equals("POST")) {
            // Get a reader over the inputstream from the request. This will
            // get a reader over the request body.
            final Reader body = new InputStreamReader(exchange.getInputStream());

            // Create a new GSON object. This is similar to creating a NewEncoder/NewDecoder in a single
            // object in Go.
            Gson gson = new Gson();

            // Decode the request body into a Task. This is essentially doing json.NewDecoder(r).Decode(&task) in Go.
            Task task = gson.fromJson(body, Task.class);
            
            try {
                task = this.store.insertNewTask(task.getTitle(), task.isComplete());
            } catch (SQLException e) {
                throw new HttpException("error inserting task: " + e.getMessage(), StatusCodes.INTERNAL_SERVER_ERROR);
            }
            exchange.getResponseHeaders().add(Headers.CONTENT_TYPE, "application/json");
            exchange.getResponseSender().send(gson.toJson(task));
        } else {
            throw new HttpException("method not allowed", StatusCodes.METHOD_NOT_ALLOWED);
        }
    }
}
