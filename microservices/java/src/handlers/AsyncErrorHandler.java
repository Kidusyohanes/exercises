package handlers;

import io.undertow.server.HttpHandler;
import io.undertow.server.HttpServerExchange;

public class AsyncErrorHandler extends ErrorHandler {

    public AsyncErrorHandler(HttpHandler handler) {
        super(handler);
    }

    public void handleRequest(HttpServerExchange exchange) {
        if (exchange.isInIoThread()) {
            exchange.dispatch(this);
            return;
        }
        exchange.startBlocking();
        super.handleRequest(exchange);
    }
}
