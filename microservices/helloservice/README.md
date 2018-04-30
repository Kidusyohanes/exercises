# Example Node.js Microservice

- At the command line, execute `npm init -y` in this directory to start a Node.js project.
- Execute `npm install --save express` to install the express web server framework.
- Create an `index.js` file for your new web server, and start with this code:

```javascript
//set interpreter into strict mode
"use strict";
//load the express module from the node_modules directory
const express = require("express");
//create a new express web application
const app = express();
```

See the [Node.js Microservices tutorial](https://drstearns.github.io/tutorials/nodeweb/) for more details.