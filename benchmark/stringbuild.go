package benchmark

import (
	"strings"
)

//Concat concatenates all the strings passed into it
//using simple string concatenation
func Concat(parts ...string) string {
	ret := ""
	for _, s := range parts {
		ret += s
	}
	return ret
}

//Concat2 concatenates all the strings passed into it
//using strings.Builder
func Concat2(parts ...string) string {
	builder := strings.Builder{}
	for _, s := range parts {
		builder.WriteString(s)
	}
	return builder.String()
}
