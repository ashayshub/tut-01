package main

import (
	"log"
	"net/http"
	"time"
)

type customHandler struct{}

func (c customHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	resp := []byte("Hello World")
	_, err := rw.Write(resp)
	if err != nil {
		log.Fatal("Something went wrong")
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
