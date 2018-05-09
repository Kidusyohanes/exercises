package models;

public class Task {

    private int id;
    private String title;
    private boolean complete;

    // There is no modifier here, such as "public" since there is no
    // need for another package to create their own Tasks.
    // See https://docs.oracle.com/javase/tutorial/java/javaOO/accesscontrol.html
    Task(int id,  String name, boolean complete) {
        this.id = id;
        this.title = name;
        this.complete = complete;
    }

    public int getId() {
        return id;
    }

    public String getTitle() {
        return title;
    }

    public boolean isComplete() {
        return complete;
    }
}
