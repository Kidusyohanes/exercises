# Common Bugs In 344 Assignments

## All Assignments
1. Not returning after reporting errors to the user. `http.Error` does not stop program execution.
2. Returning the improper http code.
3. Resource leaks. Aka Forgetting to close persistent connections to things (like streams, and database connections).
4. Not checking the error _before_ deferring a .close(). This can, and will panic.
5. No `Access-Control-Allow-Origin: *` header.
6. Incorrect `content-type` header. hint it's: `application/json; utf-8`.

## Assignment 1
1. Not checking the http response code from the page you are scraping. (If you get an error, you need to return an error to the user!)
2. Continuing program execution when your program encounters an html.ErrorToken.
3. Parsing the whole DOM. Hint: you just need to parse the `<head>STUFF</head>`.

## Assignment 2
1.
2.
3.
4.
