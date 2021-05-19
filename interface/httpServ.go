package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type customHandler struct{}

func (c customHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprint(w, "Hello World\n")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	myHandler := customHandler{}
	s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Print("Starting http server on 8080...")
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
