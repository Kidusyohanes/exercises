package main

import (
	"log"
	"os"
)

const maxConnRetries = 5

func reqEnv(name string) string {
	val := os.Getenv(name)
	if len(val) == 0 {
		log.Fatalf("please set the %s environment variable", name)
	}
	return val
}

func main() {

}
