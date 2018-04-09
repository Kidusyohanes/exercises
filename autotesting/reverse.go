package autotesting

//Reverse returns the reverse of the string passed as `s`.
//For example, if `s` is "abcd" this will return "dcba".
func Reverse(s string) string {
	//BUG: there are a couple of nasty bugs in here.
	//Let's write some automated tests to uncover them,
	//and fix the code until the tests pass.

	chars := []byte(s)
	for i, j := 0, len(chars)-1; i < j; i = i + 1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
