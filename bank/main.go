package main


import (
	"os"
	_ "github.com/go-redis/redis"
	"net/http"
	"log"
)

/**
	HINT: Create a shared context defintion here. 
*/

func main() {
	redisAddr := os.Getenv("REDIS")
	addr := os.Getenv("ADDR")

	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	if addr == "" {
		addr = "localhost:4000"
	}

	/* Hint: use this snippet to initialize a Redis client
	client := redis.NewClient(&redis.Options{
			Addr: redisAddr,
	})
	*/

	mux := http.NewServeMux()

	mux.HandleFunc("/balance", BalanceHandler)

	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}


func BalanceHandler(w http.ResponseWriter, r *http.Request) {
	/**
		Implement BalanceHandler

		This is where we will need to implement:
			- Handling GET /balance
			- Handling PUT /balance, and its CORS pre-flight request
	*/

	if r.Method == "GET" {
		// handle GET

	} else if r.Method == "PUT" {
		// handle PUT

	} else if r.Method == "OPTIONS" {
		// handle CORS pre-flight

	} else {


	}
}
