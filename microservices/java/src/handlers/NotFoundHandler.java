package handlers;

import io.undertow.server.HttpHandler;
import io.undertow.server.HttpServerExchange;
import io.undertow.util.Headers;
import io.undertow.util.StatusCodes;

// Undertow's default behavior on a route that wasn't found is to send an empty
// response body, resulting in kind of a nasty looking message from our browsers.
// Attaching this handler as a not found handler will mimic the behavior that Go's
// webserver has when a resource is not found.
public class NotFoundHandler implements HttpHandler {
    public void handleRequest(HttpServerExchange exchange) {
        exchange.setStatusCode(StatusCodes.NOT_FOUND);
        exchange.getResponseHeaders().add(Headers.CONTENT_TYPE, "text/plain");
        exchange.getResponseSender().send("404 not found");
    }
}
