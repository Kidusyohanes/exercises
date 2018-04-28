"use strict";

//SQL statements we will need to execute
const SQL_SELECT_ALL = "select id,url,comment,votes from links";
const SQL_SELECT_BY_ID = SQL_SELECT_ALL + " where id=?";
const SQL_INSERT = "insert into links (url,comment) values (?,?)";
const SQL_UPVOTE = "update links set votes=votes+1 where id=?";

//TODO: use `npm init -y` to create a package.json file
//and then `npm install --save express mysql` to install the 
//express and mysql packages to your node_modules directory
