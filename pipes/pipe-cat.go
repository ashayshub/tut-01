// Twice io.Copy
package main

import (
	"io"
	"log"
	"os"
	"sync"
)

func ReadFile(filePath string, pw *io.PipeWriter, wg *sync.WaitGroup) {
	defer wg.Done()
	defer pw.Close()
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Could not open file:", filePath, err)
	}
	n, err := io.Copy(pw, f)
	if err != nil {
		log.Fatalf("Could not read file: %v, Error: %v, Bytes written: %v", filePath, err, n)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Improper arg count")
	}

	var wg *sync.WaitGroup = &sync.WaitGroup{}
	defer wg.Wait()

	pr, pw := io.Pipe()

	filePath := os.Args[1]
	wg.Add(1)
	go ReadFile(filePath, pw, wg)

	n, err := io.Copy(os.Stdout, pr)

	if err != nil {
		log.Fatalf("Error reading from pipe, Error: %v, bytes written: %v", err, n)
	}
}
