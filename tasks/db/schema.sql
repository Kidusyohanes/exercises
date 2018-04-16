-- TODO: create a table to store tasks
-- as defined in the models/tasks/task.go struct.
-- Use an auto-incrementing primary key for the ID field
-- see https://dev.mysql.com/doc/refman/5.7/en/create-table.html
create table tasks (
    id int primary key auto_increment not null,
    title varchar(255) not null,
    completed bool not null default false
);
