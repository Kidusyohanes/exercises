package models;

import com.mysql.jdbc.jdbc2.optional.MysqlDataSource;

import java.sql.*;

public class SQLStore {

    // Reference to the direct connection to the Database
    // that this instance of the SQLStore will use.
    private Connection conn;

    public SQLStore(String addr, String user, String password, String db) throws IllegalArgumentException, SQLException {
        if (addr.length() == 0 || user.length() == 0 || password.length() == 0 || db.length() == 0) {
            throw new IllegalArgumentException("addr, user, password, and db must all be provided");
        }

        // Parse the addr string into a host and a port.
        String[] splitAddr = addr.split(":");
        if (splitAddr.length != 2) {
            throw new IllegalArgumentException("DB address must be formatted as host:port");
        }

        String host = splitAddr[0];
        int port;
        try {
            port = Integer.parseInt(splitAddr[1]);
        } catch (Exception e) {
            throw new IllegalArgumentException("DB port is not a valid number");
        }

        // Build the DSN. Doing this with the MysqlDataSource will
        // prevent some headaches.
        MysqlDataSource dsn = new MysqlDataSource();
        dsn.setUser(user);
        dsn.setDatabaseName(db);
        dsn.setPassword(password);
        dsn.setServerName(host);
        dsn.setPort(port);

       this.conn =  dsn.getConnection();
    }

    public Task insertNewTask(String name, boolean complete) throws IllegalArgumentException, SQLException {
        if (name.length() == 0) {
            throw new IllegalArgumentException("task name must exist");
        }

        String rawQuery = "INSERT INTO tasks (title, completed) VALUES (?, ?)";
        PreparedStatement stmt = this.conn.prepareStatement(rawQuery, Statement.RETURN_GENERATED_KEYS);

        // Assign that the ? are going to be. Indexes here start at 1.
        stmt.setString(1, name);
        stmt.setBoolean(2, complete);

        // Execute the statement.
        stmt.execute();

        int id = 0;
        // Get the generated keys from the statement. We are able to do this because
        // we passed in Statement.RETURN_GENERATED_KEYS to the prepareStatement method
        // above.
        ResultSet rs = stmt.getGeneratedKeys();
        if (rs.next()) {
            id = rs.getInt(1);
        }

        return new Task(id, name, complete);
    }
}
