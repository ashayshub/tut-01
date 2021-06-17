package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func completeProcessing(c chan bool, wg *sync.WaitGroup) {
	time.Sleep(10 * time.Second)
	wg.Done()
	c <- true
}

type H struct{}

func (h H) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// setup request processing

	var c chan bool

	// setup waitgroup
	var wg sync.WaitGroup = sync.WaitGroup{}
	defer wg.Wait()

	// setup request context
	log.Printf("%v %v", r.Method, r.URL)
	ctx := r.Context()

	// set processing
	wg.Add(1)
	go completeProcessing(c, &wg)

	// wait on whichever goroutine givesup first
	select {

	case <-c:
		log.Print("Request complete.")
		fmt.Fprint(w, "Hello World\n")

	case <-ctx.Done():
		log.Print("Request cancelled.")
		fmt.Fprint(w, "Cancelled.\n")
	}

}

func main() {
	var h H = H{}
	s := &http.Server{
		Addr:           ":8080",
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
