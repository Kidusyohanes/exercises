create table links (
    id int primary key auto_increment,
    url varchar(2048) not null,
    comment varchar(1024),
    votes int not null default 0
);
