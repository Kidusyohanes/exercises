package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const usage = `
usage:
	crawl <starting-url> [num-workers=1]

If not specified, the number of workers will
default to 1. To specify more workers, pass
an integer as the second command-line argument.
`
const defaultNumWorkers = 1

type crawlResults struct {
	pageInfo *PageInfo
	err      error
}

//TODO: define a worker function

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	numWorkers := defaultNumWorkers
	if len(os.Args) > 2 {
		var err error
		numWorkers, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("number of workers must be a number: %v", err)
		}
	}
	log.Printf("starting %d workers...", numWorkers)

	//TODO: create channels for the URLs to crawl (jobs)
	//and the crawlResults of each crawl (results)

	//TODO: start workers on their own goroutines,
	//passing those channels as arguments

	//TODO: create a map to track all the URLs
	//we've already crawled so that we don't
	//crawl them multiple times

	//TODO: write the first command-line arg
	//to the URLs-to-crawl channel, and add it
	//to the map of already-crawled URLs

	//TODO: range over the results channel,
	//processing each link in the crawled page info

}
