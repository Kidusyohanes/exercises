package main

import (
	"time"
)

type Broken struct {
	test map[string]string
}

func main() {

	b := Broken{}
	go func() {
		b.test["bork"] = "0"
	}()

	time.Sleep(1000000)
}
