package middleware

//TODO: implement a Logger middleware handler
//that also captures and logs the response statusCode

// https://ndersson.me/post/capturing_status_code_in_net_http/
// https://travix.io/type-embedding-in-go-ba40dd4264df

import (
  "log"
  "net/http"
  "time"
)


type StatusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (sr *StatusRecorder) WriteHeader(statusCode int) {
	sr.statusCode = statusCode
	sr.ResponseWriter.WriteHeader(statusCode)
}

type Logger struct {
	wrappedHandler http.Handler
}

func NewLogger(handlerToWrap http.Handler) http.Handler {
	return &Logger{handlerToWrap}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	begin := time.Now()
	statusRec := &StatusRecorder{w}
	l.wrappedHandler.ServeHTTP(statusRec, r)
	log.Printf("%s %s - %d %v", r.Method, r.URL.Path, statusRec.statusCode, time.Since(begin))
}
