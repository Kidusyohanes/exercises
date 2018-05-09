package handlers;

import io.undertow.server.HttpHandler;
import io.undertow.server.HttpServerExchange;

public class AsyncErrorHandler extends ErrorHandler {

    public AsyncErrorHandler(HttpHandler handler) {
        super(handler);
    }

    // The normal error handler doesn't work for async operations
    // since the try-catch logic will end up on a different thread
    // then the logic that is being wrapped. To handle this we will
    // dispatch the error handling logic as well so then it properly
    // handles the try-catch and any exceptions that are thrown.
    public void handleRequest(HttpServerExchange exchange) {
        if (exchange.isInIoThread()) {
            exchange.dispatch(this);
            return;
        }
        exchange.startBlocking();
        super.handleRequest(exchange);
    }
}
