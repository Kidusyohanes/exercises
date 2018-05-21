package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
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

func reportResults(r *crawlResults, results chan<- *crawlResults) {
	results <- r
}

//TODO: define a worker function
func worker(urlsToFetch <-chan string, results chan<- *crawlResults) {
	log.Printf("worker starting...")
	for URL := range urlsToFetch {
		log.Printf("crawling %s...", URL)
		info, err := GetPageInfo(URL)
		go reportResults(&crawlResults{info, err}, results)
		time.Sleep(time.Second)
	}
}

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
	urlsToFetch := make(chan string, 2048)
	results := make(chan *crawlResults, 2048)
	seen := map[string]bool{}

	//TODO: start workers on their own goroutines,
	//passing those channels as arguments
	for i := 0; i < numWorkers; i++ {
		go worker(urlsToFetch, results)
	}

	//TODO: write the first command-line arg
	//to the URLs-to-crawl channel, and add it
	//to the map of already-crawled URLs
	startingURL := os.Args[1]
	seen[startingURL] = true
	urlsToFetch <- startingURL
	urlsOutstanding := 1

	//TODO: range over the results channel,
	//processing each link in the crawled page info
	for cr := range results {
		urlsOutstanding--
		if cr.err != nil {
			log.Printf("error crawing: %v", cr.err)
			continue
		}
		for _, link := range cr.pageInfo.Links {
			if !seen[link] {
				seen[link] = true
				if ShouldCrawl(link) {
					log.Printf("adding %s to work queue", link)
					urlsToFetch <- link
					urlsOutstanding++
				}
			}
		}
		log.Printf("%d URLs outstanding", urlsOutstanding)
		if urlsOutstanding == 0 {
			break
		}
	}
}
