package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func sucessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received!")

	w.WriteHeader(http.StatusOK)
	delay := r.URL.Query().Get("delay")

	fmt.Println(delay)

	if delay != "" {
		delayDuration, err := strconv.Atoi(delay)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error!"))
			return
		}

		time.Sleep(time.Second * time.Duration(delayDuration))
	}

	w.Write([]byte("Hello World!"))
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error!"))
}

func main() {
	http.HandleFunc("/sucess", sucessHandler)
	http.HandleFunc("/error", errorHandler)
	http.ListenAndServe(":8080", nil)
}
