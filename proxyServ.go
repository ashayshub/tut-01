package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {

	target, err := url.Parse("http://localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	rp := httputil.NewSingleHostReverseProxy(target)

	proxy := &http.Server{
		Addr:           ":8081",
		Handler:        rp,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Print("Starting proxy server on 8081, proxying to localhost:8080...")
	err2 := proxy.ListenAndServe()

	if err2 != nil {
		log.Fatal(err2)
	}
}
