# Succeeding in INFO 344
INFO 344 is a thrilling exploration into the world of server-side web development. You will learn how many of the processes and designs that are used in industry work and how to apply them to projects that you undertake as part of an internshop or full-time job, as well as any of your fun side projects.

However, as fun and valuable as this course is it is also **demanding** and we will expect much from you. You should expect to put in considerable effort and time into this course. It is not uncommon to be intimidated by this course, but don't be affraid to push through it! We have confidence that each and every one of you can succeed and flourish! It is not our intention to scare you away with the wording in this document, we only want to help set you up for success.

We have compiled this document to help you know what you should expect going into this course, as well as to help avoid wasting time or losing unnecessary points.

# Helpful Tips
1. Go to lecture - every time. Seriously.
2. Ask questions! If you have a question, other people probably do too. We are here to help you, and we know that you can succeed!
3. Go to office hours, and ask more questions on Slack.
4. Learn how to Google effectively -- somebody else has probably had the same problem as you.
4.1. Read the questions before jumping to answers on Stack Overflow, it's really easy to forget to learn while you are trying to solve a problem. 
5. Learn how to isolate your bugs -- what happened? what did you expect to happen? when was it last working?
5.1. Rubber duck debugging.
6. Do the readings! Read the docs, just read all the time.
7. git commit -m "often"
8. Write tests!
9. Write terraform configurations / deploy scripts. -- at least write local dev scripts. This will seriously save you so many headaches.
10. Learn how to test properly. #1 test in docker locally #2 test in docker remotely
11. Familiarize yourself with a debugger. Some IDEs such as GoLand come with these built in, but many don't. A popular one is the [Delve Debugger](https://github.com/derekparker/delve/tree/master/Documentation/installation), which can be hooked up to VS Code.

## How INFO 344 Differs From Other Coding Courses
1. **This is a much more challenging course:** the range of topics, scope of assignments, and breadth of knowledge you're going to learn are all much more than many coding courses you've taken.
2. **This is NOT a weedout course:** we are not looking for ways to make you fail. In fact, we want everyone to succeed and we will do everything we can to make that happen. But we will expect the same effort from you.
3. **The assignment specs are all essentially minimum requirement specs:** If you want to add more things to each assignment, please do! If you want to add some awesome new feature thats not described in the spec, do it! We would love to see what you can do! We have a fun policy: if you want to do cool shit, do cool shit. And if you tell us what cool shit you did in an assignment submission, we may even throw some extra credit points your way.

## Our Expectations

This is likely your 4th or 5th coding class that you have taken and we will grade your assignments as such. This means that we will expect you to employ coding practices that you've learned in previous classes even if they are not explicitly discussed in lecture, lab, or the assignment specs. 

For example:

1. Appropriate variable/function naming convention
2. Adherence to the DRY (don't repeat yourself) principle
3. Proper encapsulation
4. and anything taught in CSE 142/143, INFO 201/343, etc

Everyone on the teaching team is ready and would love to help you. However, before you ask us a question we will expect you to have done a few things first (as applicable):

1. Read Dr. Stearns' tutorial on the relevant topic
2. Researched (Googled) the problem or topic
3. Attempted to debug the problem yourself


## Common Pitfalls

1. **Pull Requests created incorrectly:** Please check that your PRs correctly compare branches. We can only grade the code that you submit in your PR, so if there is no code to grade then we cannot give you points for your implementation.
2. **API servers down:** For many assignments we will grade your assignment's functionality by running automated tests against your API server. If this server is unresponsive or down for any reason we cannot grade your functionality and therefore cannot give you points for your functionality. Please ensure that your API server does respond to requests even after multiple uses of your own test suite.
3. **Not telling us about extenuating circumstances:** Please, please, please tell us if you have issues outside of the class will affect your performance. We are not here to make your life difficult, and this is not a weed-out course so we will do what we can to accomodate you. However, we cannot do this if you do not tell us first.
4. **Not starting assignments early:** You hear this for every course you take, we get it. We hear it too. Please take this warning from student-to-student: start early! These assignments can be quite extensive and there will likely be many issues you will run into. Expect to take a decent amount of time on each one.
5. **Not cleaning up code before submitting:** We know that after spending 6 hours working on an assingment the last thing you want to do is clean it up. However, this is important. Make sure to remove all commented out code and unnecessary/debug print statements. We will take points off for every occurance of these. This can be pretty easy to do using many IDEs search all function.
6. **Not submitting assignments correctly:** Although each assignment requires generally the same thing (a link to the PR, a link to the deployed API server, and a link to the deployed web client), please double check what the assignment asks you to submit. If any of these are missing and we have to find them ourselves we will take points off.

## Common Bugs In Assignment 1
1. Not returning after reporting errors to the user. `http.Error` does not stop program execution.
2. Returning the improper http code.
3. Resource leaks. Aka Forgetting to close persistent connections to things (like streams, and database connections).
4. Not checking the error _before_ deferring a .close(). This can, and will panic.
5. Not checking the http response code from the page you are scraping. (If you get an error, you need to return an error to the user!)
6. No `Access-Control-Allow-Origin: *` header.
7. Incorrect `content-type` header. hint it's: `application/json; utf-8`.
8. Continuing program execution when your program encounters an html.ErrorToken.
9. Parsing the whole DOM. Hint: you just need to parse the `<head>STUFF</head>`.

# Other Notes

We encourage you to try JetBrains' [GoLand](https://www.jetbrains.com/go/) (Ethan recommends you use Vim). Goland is JetBrains' industry IDE built specifically for working with Golang, and they offer free student licenses! If you have used any JetBrains IDE before then you know they are quite robust and offer a suite of neat features. Among the ones that may help you in this course are the built-in debugger and the ability to connect to SQL databases which provided SQL statement auto completion when witing queries in Go.

However, GoLand can be intimidating when you first use it but will be well worth it. Also note that it will require a few extra steps to set up, but these are not cumbersome and Brendan would be happy to help you complete these.

The most important thing is to use an IDE/text editor that you are comfortable with. There are many great options and we will not require you to use any specific one, so feel free to make your own choice.
