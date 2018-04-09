# Delve
Using a debugger is one of the things that separates amateur and professional developers, and though it will seem harder than print statements at first, it will serve you well for hard bugs!
[Installation](https://github.com/derekparker/delve/tree/master/Documentation/installation)
macOS: `xcode-select --install` `go get -u github.com/derekparker/delve/cmd/dlv`

## VS Code Setup
Open up vscode and make sure that the Go extension is installed (this should have happened automatically). The Go extension now installes delve as of 0.6.56.

## Exercise
`factorial.go` is a program that calculates the factorial of a number. It has bugs, so let's use Delve to find the bugs!

### Start Debugging in VS Code
https://code.visualstudio.com/docs/editor/debugging
Enter the Debug Pane by clicking the debug icon on the left side of vscode.
![Debug Pane](https://code.visualstudio.com/assets/docs/editor/debugging/debugicon.png)
Set a breakpoint at the start of the `factorial()` function by clicking to the left of line 5. A red dot should show up. Now click the green play button at the top left.
As you step in, and step over the function, the current state of the program's variables are showin the the `VARIABLES` pane.

### Start Debugging On The Command Line
```
$ dlv debug
Type 'help' for list of commands.
(dlv) break main.factorial
Breakpoint 1 set at 0x10b1380 for main.factorial() ./factorial.go:5
(dlv) continue
> main.factorial() ./factorial.go:5 (hits goroutine(1):1 total:1) (PC: 0x10b1380)
     1:	package main
     2:
     3:	import "fmt"
     4:
=>   5:	func factorial(num int) int {
     6:		var i, j int
     7:		for i = 1; i < num; i++ {
     8:			j = j * i
     9:		}
    10:		return j
(dlv)
```
### Common Commands
`break` alias `b`: Sets a breakpoint. See documentation, usually set with <filename>.<function> or <filename>.<line> or just <line>  
`continue` alias `c`: Run until breakpoint or program termination.  
`step` alias `s`: Single step through the program. Known as step into in vscode.
`next` alias `n`: Step over to next source line. (This will _not_ step into function calls). Known as step over in vscode.
`print` alias `p`: Evaluate an expression. (This can be used to print variables or conditions like: `print i`)  
`locals`: Print local variables.  
