package handlers;

import io.undertow.io.Sender;
import io.undertow.server.HttpHandler;
import io.undertow.server.HttpServerExchange;
import io.undertow.util.Headers;
import io.undertow.util.StatusCodes;

import java.util.Map;
import java.util.Deque;

public class HelloHandler implements HttpHandler {

    // The only method that needs to be implemented to satisfy the Undertow.HttpHandler
    // interface is this handleRequest method. This is the Undertow equivelant of the
    // http.ServeHTTP function in Go
    public void handleRequest(HttpServerExchange exchange) {
        exchange.getResponseHeaders().add(Headers.CONTENT_TYPE, "text/plain");
        Sender rs = exchange.getResponseSender();

        Map<String, Deque<String>> queries = exchange.getQueryParameters();
        Deque<String> names = queries.get("name");
        if (names != null && names.peek().length() > 0) {
            String name = names.pop();
            rs.send("hello " + name + "!");
        } else {
            throw new HttpException("name query string param not provided", StatusCodes.BAD_REQUEST);
        }
    }
}
