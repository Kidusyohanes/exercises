/**
 * UnsetVariableException defines an exception to be used
 * in cases where environment variables which are required
 * do not have values.
 */
class UnsetVariableException extends RuntimeException {

    // UnsetVariableException creates a new UnsetVariableException
    // with the given message.
    UnsetVariableException(String message) {
        super(message);
    }
}
