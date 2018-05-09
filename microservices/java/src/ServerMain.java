import handlers.*;
import io.undertow.Undertow;
import io.undertow.Undertow.*;
import io.undertow.server.handlers.PathHandler;
import models.SQLStore;

import java.sql.SQLException;

import static io.undertow.Handlers.path;

public class ServerMain {
    public static void main(String[] args) throws SQLException {

        // Get required server env vars
        int port = Integer.parseInt(requireENV("JAVA_PORT", "80"));
        String addr = requireENV("JAVA_ADDR", "0.0.0.0");

        // Get required MySQL env vars
        String mysqlAddr = requireENV("JAVA_MYSQL_ADDR", "");
        String mysqlDB = requireENV("JAVA_MYSQL_DB", "");
        String mysqlPW = requireENV("JAVA_MYSQL_PASS", "");
        String mysqlUser = requireENV("JAVA_MYSQL_USER", "root");

        // Create a new SQLStore object that we can give to our handlers
        // that need to connect to a database.
        SQLStore store = new SQLStore(mysqlAddr, mysqlUser, mysqlPW, mysqlDB);

        // Create the server builder object.
        // We will use this to specify properties on our
        // server prior to starting it.
        Builder builder = Undertow.builder();
        builder.addHttpListener(port, addr);

        // Undertow doesn't have a "mux" per se, but has something very similar.
        // This PathHandler will function more or less as an http.ServeMux from Go.
        PathHandler mux = path();
        mux.addPrefixPath("/", new NotFoundHandler());
        mux.addPrefixPath("/tasks", new AsyncErrorHandler(new TasksHandler(store)));
        mux.addPrefixPath("/hello", new ErrorHandler(new HelloHandler()));
        mux.addPrefixPath("/time", new TimeHandler());

        // Tell the server builder to use the PathHandler we defined above.
        builder.setHandler(mux);

        //  Build and start the server.
        Undertow server = builder.build();
        System.out.printf("Server is listening on %s:%d...\n", addr, port);
        server.start();
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
