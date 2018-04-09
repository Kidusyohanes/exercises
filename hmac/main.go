package main

import (
	"errors"
	"fmt"
	"os"
)

//usage is the usage string shown when not enough
//arguments are passed to this command
const usage = `
usage:
	hmac sign secretKey < file-to-sign
	hmac verify secretKey signature < file-that-was-signed

When using 'sign' the output will be a base64-encoded
HMAC digital signature for the file using the secretKey.

When using 'verify' the output will be "signature is valid"
if the provided signature is valid for the file and secretKey,
or "signature is INVALID" if the provided signature is
invalid for the file and secretKey
`

//showErrorAndExit prints the error message
//and exits with a non-success code
func showErrorAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

//showUsageAndExit shows the usage string
//and exits with a non-success code
func showUsageAndExit() {
	showErrorAndExit(errors.New(usage))
}

//main is the main entry point for this command
func main() {
	//TODO: implement this function so that after you
	//run `go install`, your program will support the two
	//usage examples noted above in the usage string.
	//Note that os.Stdin is an io.Reader for the stdin stream.

	//for example, running this command:
	//  hmac sign mysecretkey < test.txt
	//should print the following *base64-encoded* HMAC
	//signature to the terminal:
	// AHYFAcm0TnHvpWdQoyeWdeHgy-t54nK-4u8xsK2_cTg=

	//and running this command:
	//  hmac verify mysecretkey AHYFAcm0TnHvpWdQoyeWdeHgy-t54nK-4u8xsK2_cTg= < test.txt
	//should print "signature is valid" is to the terminal
	//while running this command (signature has been modified):
	//  hmac verify mysecretkey BHYFAcm0TnHvpWdQoyeWdeHgy-t54nK-4u8xsK2_cTg= < test.txt
	//should print "signature is INVALID" to the terminal.
	//similarly, running this command (key has been modified):
	//  hmac verify differentkey AHYFAcm0TnHvpWdQoyeWdeHgy-t54nK-4u8xsK2_cTg= < test.txt
	//should also print "signature is INVALID" to the terminal.

	//Pro Tip: use shell variables to hold your key and
	//the signature returned from `hmac sign` so that you
	//can more easily pass them to the commands:
	//
	//  #set SIGNKEY to your secret signing key
	//  SIGNKEY=mysecretkey
	//  # set SIG equal to the output of hmac sign
	//  SIG=$(hmac sign $SIGNKEY < test.txt)
	//  # to see the value of SIG...
	//  echo $SIG
	//  # pass it to hmac verify
	//  hmac verify $SIGNKEY $SIG < test.txt
	//
	//you can also use a shell variable to hold your key!

}
