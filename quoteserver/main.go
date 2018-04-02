package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const headerContentType = "Content-Type"
const headerAccessControlAllowOrigin = "Access-Control-Allow-Origin"
const contentTypeJSON = "application/json"
const originAll = "*"

//Quote represents a single quote record
type Quote struct {
	Quote  string `json:"quote,omitempty"`
	Author string `json:"author,omitempty"`
	Genre  string `json:"genre,omitempty"`
}

//QuoteSlice is a slice of Quotes
type QuoteSlice []*Quote

//QuoteHandler handles requests for quotes
type QuoteHandler struct {
	quotes QuoteSlice
}

//NewQuoteHandler constructs a new QuoteHandler
func NewQuoteHandler(quotes QuoteSlice) *QuoteHandler {
	return &QuoteHandler{
		quotes: quotes,
	}
}

//ServeHTTP handles HTTP requests for the QuoteHandler
func (qh *QuoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//allow cross-origin AJAX requests
	w.Header().Add(headerAccessControlAllowOrigin, originAll)
	//tell the client the response body is encoded in JSON
	w.Header().Add(headerContentType, contentTypeJSON)

	//respond with a randomly-chosen quote
	//encoded into JSON
	i := rand.Intn(len(qh.quotes))
	if err := json.NewEncoder(w).Encode(qh.quotes[i]); err != nil {
		http.Error(w, fmt.Sprintf("error encoding JSON: %v", err), http.StatusInternalServerError)
	}
}

func loadQuotes(r io.Reader, numExpected int) (QuoteSlice, error) {
	reader := csv.NewReader(r)
	reader.Comma = ';' //quotes have embedded commas, so this file uses ; as a delimeter

	//read but ignore the column headings
	_, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading headers: %v", err)
	}

	//create a QuoteSlice with capacity = numExpected
	quotes := make(QuoteSlice, 0, numExpected)

	//for each record
	for {
		//read the fields
		fields, err := reader.Read()
		//if we are at the end of the file, return the quotes
		if err == io.EOF {
			return quotes, nil
		}
		//if we got some other error, return that
		if err != nil {
			return nil, fmt.Errorf("error reading record: %v", err)
		}
		//construct a new Quote and append to the slice
		q := &Quote{
			Quote:  fields[0],
			Author: fields[1],
			Genre:  fields[2],
		}
		quotes = append(quotes, q)
	}
}

func main() {
	//open the quotes.csv file
	f, err := os.Open("quotes.csv")
	if err != nil {
		log.Fatalf("error opening quotes.csv file: %v", err)
	}

	//load the quotes
	quotes, err := loadQuotes(f, 75966)
	if err != nil {
		log.Fatalf("error parsing quotes.csv file: %v", err)
	}
	f.Close()
	log.Printf("loaded %d quotes", len(quotes))

	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//get address to listen on from ADDR environment variable
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		//default to :80 (port 80, any host)
		addr = ":80"
	}

	mux := http.NewServeMux()
	mux.Handle("/", NewQuoteHandler(quotes))

	log.Printf("server is listening at http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
