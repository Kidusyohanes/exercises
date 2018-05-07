import java.sql.SQLException;

public class ServerMain {
    public static void main(String[] args) throws SQLException {
        // Implement the main method here
    }

    /**
     * Gets the environment variable at "env" or returns the default value "def" if
     * the environment variable is not set.
     * @param env environment variable to get value of
     * @param def default value to use if environment variable is not set
     * @return String the value of the environment variable or the default value if not defined.
     * @throws UnsetVariableException If environment variable is not set and no default is provided.
     */
    private static String requireENV(String env, String def) throws UnsetVariableException {
        String envVal = System.getenv(env);
        if (envVal == null || envVal.length() == 0) {
            if (def.length() == 0) {
                throw new UnsetVariableException(env + " is not set and no default was provided");
            }
            return def;
        }
        return envVal;
    }
}
