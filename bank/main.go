package main


import (
	"os"
	"github.com/go-redis/redis"
	"net/http"
	"log"
	"strconv"
	"fmt"
	"encoding/json"
)

/**
	HINT: Create a shared context defintion here.
*/
type SharedResource struct {
	client * redis.Client
}

type Transaction struct {
	Amt string `json:"amt"`
}


func main() {
	redisAddr := os.Getenv("REDIS")
	addr := os.Getenv("ADDR")

	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	if addr == "" {
		addr = "localhost:4444"
	}

	/* Hint: use this snippet to initialize a Redis client
	*/
	client := redis.NewClient(&redis.Options{
			Addr: redisAddr,
	})

	ctx := SharedResource{client: client}

	mux := http.NewServeMux()

	mux.HandleFunc("/balance", ctx.BalanceHandler)

	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}


func fetchBalance( ctx * SharedResource ) int {
	var currBal int
	bal, err := ctx.client.Get("Balance").Result()

	if err != nil {
		bal = ""
	}

	if bal == "" {
		currBal = 0
	} else {
		currBal, _ = strconv.Atoi(bal)
	}

	return currBal
}


func (ctx *SharedResource) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	/**
		Implement BalanceHandler

		This is where we will need to implement:
			- Handling GET /balance
			- Handling PUT /balance, and its CORS pre-flight request
	*/

	if r.Method == "GET" {
		if r.Header.Get("Origin") == "" {
			http.Error(w, "Bad CORS Request", http.StatusBadRequest)
			return
		}

		// handle GET
		currBal := fetchBalance(ctx)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		payload := fmt.Sprintf("{ \"amt\": %d }", currBal)
		w.Write([]byte(payload))
	} else if r.Method == "PUT" {
		// handle PUT
		if r.Header.Get("Origin") == "" {
			http.Error(w, "Bad CORS Request", http.StatusBadRequest)
			return
		}

		// validate request
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Not valid Content-Type", http.StatusBadRequest)
			return
		}

		// parse JSON body
		decoder := json.NewDecoder(r.Body)
		var newTransaction Transaction
		err := decoder.Decode(&newTransaction)
		if err != nil {
				http.Error(w, "Request body could not be parsed", http.StatusBadRequest)
				return
		}

		// update the Balance
		currBal := fetchBalance(ctx)
		inc, _ := strconv.Atoi(newTransaction.Amt)
		currBal += inc

		// save change
		ctx.client.Set("Balance", strconv.Itoa(currBal), 0).Err()

		// If you need to expose additional headers, make sure to include "Access-Control-Expose-Header"

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		payload := fmt.Sprintf("{ \"amt\": %d }", currBal)
		w.Write([]byte(payload))
	} else if r.Method == "OPTIONS" {
		// handle CORS pre-flight

		if r.Header.Get("Access-Control-Request-Method") == "" || r.Header.Get("Origin") == "" {
				http.Error(w, "Bad CORS Pre-flight request", http.StatusBadRequest)
				return
		}

		/*
			Remember: CORS Pre-flight request MUST return 200 for it to be considered
			"successful" by the client
			Communicate to the client that PUT is an acceptible CORS request method
			Use Access-Control-Max-Age to cache the CORS pre-flight request's response
		  If you want to allow complex headers like Authorization, make sure to include "Access-Control-Allow-Header"
		*/
		w.Header().Set("Access-Control-Allow-Methods", "PUT, PATCH")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600000")
		w.WriteHeader(http.StatusOK)
	} else {


	}
}
