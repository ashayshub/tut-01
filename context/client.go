package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func RequestInfo(ctx context.Context, client *http.Client, c chan bool, wg *sync.WaitGroup) {
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/", nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	fmt.Println(resp)
	c <- true
	wg.Done()
}

func main() {

	// setup waitgroup
	var wg sync.WaitGroup = sync.WaitGroup{}
	defer wg.Wait()

	// Setup http client
	var tr *http.Transport = &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	var client *http.Client = &http.Client{Transport: tr}

	// setup context
	var ctx context.Context = context.Background()
	var d chan bool

	// setup deadline
	cancelDeadline := time.Now().Local().Add(time.Second * 5)
	ctx, cancel := context.WithDeadline(ctx, cancelDeadline)

	// execute request
	wg.Add(1)
	go RequestInfo(ctx, client, d, &wg)

	// wait on signal gives up first
	select {
	case <-ctx.Done():
		cancel()
		fmt.Println("Context done")
	case <-d:
		fmt.Println("Request Complete")
	}
	fmt.Println("Done.")
}
