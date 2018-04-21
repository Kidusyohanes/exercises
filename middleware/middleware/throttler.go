package middleware

import (
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

const headerRetryAfter = "Retry-After"

//Throttler is a request throttling handler
type Throttler struct {
	wrappedHandler http.Handler
	redisClient    *redis.Client
	maxRequests    int64
	duration       time.Duration
}

//Throttle wraps handlerToWrap with a Throttler handler. The redisClient parameter should be
//a connected redis database where the throttler can track number of requests per client.
//The maxRequests and duration parameters should be set to the maximum number of requests
//that can be made by a given client within the duration.
func Throttle(handlerToWrap http.Handler, redisClient *redis.Client, maxRequests int64, duration time.Duration) http.Handler {
	return &Throttler{
		wrappedHandler: handlerToWrap,
		redisClient:    redisClient,
		maxRequests:    maxRequests,
		duration:       duration,
	}
}

//getClientKey returns a unique key for the client
//that made the HTTP request. This key can be used
//to track the number of requests made by a given client.
func getClientKey(r *http.Request) string {
	//if the request contains an X-Forwarded-For header,
	//use the first address in that list.
	//https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For
	//otherwise, use r.RemoteAddr.
	panic("implement this function")
}

//ServeHTTP handles a request sent to the Throttle handler
func (t *Throttler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO: implement this function to track the
	//number of requests made by each distinct client
	//within the t.duration. If the number of requests exceeds
	//the t.maxRequests value within the duration,
	//respond with an http.StatusTooManyRequests error.
	//Otherwise, call .ServeHTTP() on the wrapped handler.
}
