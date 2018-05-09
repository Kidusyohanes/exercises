package handlers;

import io.undertow.server.HttpHandler;
import io.undertow.server.HttpServerExchange;
import io.undertow.util.Headers;

import java.util.Date;

public class TimeHandler implements HttpHandler {
    public void handleRequest(HttpServerExchange exchange) {
        Date date = new Date();
        exchange.getResponseHeaders().add(Headers.CONTENT_TYPE, "text/plain");
        exchange.getResponseSender().send(date.toString());
    }
}
