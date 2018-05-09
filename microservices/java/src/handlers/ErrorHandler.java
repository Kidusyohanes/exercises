package handlers;

import io.undertow.server.HttpHandler;
import io.undertow.server.HttpServerExchange;
import io.undertow.util.StatusCodes;

public class ErrorHandler implements HttpHandler {

    // Handler that this error handler is wrapping.
    private HttpHandler handler;

    public ErrorHandler() {
        throw new IllegalArgumentException("a handler to wrap must be provided");
    }

    public ErrorHandler(HttpHandler handler) {
        if (handler == null) {
            throw new IllegalArgumentException("wrapped handler cannot be null");
        }
        this.handler = handler;
    }

    public void handleRequest(HttpServerExchange exchange) {
        try {
            this.handler.handleRequest(exchange);
        } catch (HttpException e) {
            exchange.setStatusCode(e.getStatus());
            exchange.getResponseSender().send(e.getMessage());
        } catch (Exception e) {
            exchange.setStatusCode(StatusCodes.INTERNAL_SERVER_ERROR);
            exchange.getResponseSender().send(e.getMessage());
        }
    }
}
