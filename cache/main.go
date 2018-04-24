package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func worker(c *TTLCache) {
	//start a never ending loop...
	for {
		//...that sets and gets random keys/values
		//from the shared TTLCache
		i := rand.Intn(20)
		k := fmt.Sprintf("%d", i)
		log.Printf("setting %s=%d", k, i)
		c.Set(k, i, time.Second*5)
		i2 := c.Get(k).(int)
		log.Printf("got %d", i2)
		time.Sleep(time.Millisecond * time.Duration(i))
	}
}

func main() {
	//seed the pseudo-random number generator
	rand.Seed(time.Now().UnixNano())

	//create a new shared TTLCache
	c := NewTTLCache()
	//start 10 workers on separate goroutines
	//each worker will be running concurrently
	//with the others
	for i := 0; i < 10; i++ {
		//the `go` keyword before a function call
		//will start that function on a new goroutine
		//and continue executing the rest of this code
		//concurrently
		go worker(c)
	}

	//since the workers are running on separate
	//goroutines from the main one, and since
	//a go program will exit after the main()
	//function ends, use time.Sleep() to wait
	//for a while before exiting so that we can
	//see the unprotected cache fail
	time.Sleep(time.Hour)
}
