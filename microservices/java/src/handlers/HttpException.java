package handlers;

// HttpException allows us to set an HTTP status code when throwing
// an exception so the ErrorHandler that we use to wrap the handler
// that throws this exception knows what status code to send to the
// client.s
public class HttpException extends RuntimeException {

    private int status;

    public HttpException(String message, int status) {
        super(message);
        this.status = status;
    }

    public int getStatus() {
        return status;
    }
}
