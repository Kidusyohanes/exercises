package main

import (
		"log"
    "net/http"
    "os"
		"fmt"
		"strconv"
)

func main() {
    //get the value of the ADDR environment variable
    addr := os.Getenv("ADDR")

    if len(addr) == 0 {
        addr = ":5003"
    }


    mux := http.NewServeMux()

    mux.HandleFunc("/", CookieHandler)
		mux.HandleFunc("/api", ApiHandler)

    log.Printf("server is listening at %s...", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
}

/*
	TODO:
	1. Create a Redis client in the main function.
	2. Use the Redis client to keep count, instead of using cookies
*/



func ApiHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		cookies := r.Cookies()
		fmt.Printf("%v", cookies)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("GOTCHA! Transfered Money to Account"))
	} else {
		http.Error(w, "Bad Request", 400)
	}
}


func CookieHandler(w http.ResponseWriter, r *http.Request) {

	cookies := r.Cookies()

	fmt.Printf("%v", cookies)

	// look for visit_count in cookies
	// note, cookies are strings
	var visit_count string;

	for _, cookie := range cookies {
		 fmt.Println(cookie.Name)
		 if cookie.Name == "visit_count" {
			 visit_count = cookie.Value
			 break
		 }
	}

	if visit_count == "" {
		visit_count = "0"
	}

	// increment visit_count
	visit_count_num, _ := strconv.Atoi(visit_count)


	// increment if index
	if r.URL.Path == "/" {
		visit_count_num++
	}

	// convert it back
	visit_count = strconv.Itoa(visit_count_num)

	fmt.Printf("%v\n", visit_count)

	http.SetCookie(w, &http.Cookie{
		Name: "visit_count",
		Value: visit_count,
		Path: "/",
		HttpOnly: true,
	})

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("Your visit count is: "))
	w.Write([]byte(visit_count))
}
