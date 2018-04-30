"use strict";

const express = require("express");

const app = express();
const addr = process.env.ADDR || ":80";
const [host, port] = addr.split(":");
const portNum = parseInt(port)
if (isNaN(portNum)) {
    throw new Error("port number is not a number");
}

app.get("/hello", (req, res, next) => {
    let userJSON = req.get("X-User");
    if (userJSON) {
        let user = JSON.parse(userJSON);
        res.json(user);
    } else {
        res.status(401).send("please sign in");
    }
});

app.listen(portNum, host, () => {
    console.log(`server is listening at http://${addr}`)
});