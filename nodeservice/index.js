"use strict";

//SQL statements we will need to execute
const SQL_SELECT_ALL = "select id,url,comment,votes from links";
const SQL_SELECT_BY_ID = SQL_SELECT_ALL + " where id=?";
const SQL_INSERT = "insert into links (url,comment) values (?,?)";
const SQL_UPVOTE = "update links set votes=votes+1 where id=?";

//TODO: use `npm init -y` to create a package.json file
//and then `npm install --save express mysql` to install the 
//express and mysql packages to your node_modules directory
const express = require("express");
const mysql = require("mysql");

const app = express();

let db = mysql.createPool({
    host: process.env.MYSQL_ADDR,
    database: process.env.MYSQL_DATABASE,
    user: "root",
    password: process.env.MYSQL_ROOT_PASSWORD
});

//add JSON-parsing middleware to the app
app.use(express.json());

app.get("/links", (req, res, next) => {
    db.query(SQL_SELECT_ALL, (err, rows) => {
        if (err) {
            return next(err);
        }
        res.json(rows);
    });
});

app.post("/links", (req, res, next) => {
    db.query(SQL_INSERT, [req.body.url, req.body.comment], (err, results) => {
        if (err) {
            return next(err);
        }
        let newID = results.insertId;
        db.query(SQL_SELECT_BY_ID, [newID], (err, rows) => {
            if (err) {
                return next(err);
            }
            res.json(rows[0]);
        });
    });
});

app.post("/links/:linkID/votes", (req, res, next) => {
    //the ID portion of the resource path is available at:
    // req.params.linkID
    db.query(SQL_UPVOTE, [req.params.linkID], (err, results) => {
        if (err) {
            return next(err);
        }
        db.query(SQL_SELECT_BY_ID, [req.params.linkID], (err, rows) => {
            if (err) {
                return next(err);
            }
            res.json(rows[0]);
        });
    });
});

app.use((err, req, res, next) => {
    if (err.stack) {
        console.error(err.stack);
    }
    res.status(500).send("Something bad happened. Sorry.");
});

app.listen(4000, "127.0.0.1", () => {
    console.log(`server is listening at http://localhost:4000`);
})