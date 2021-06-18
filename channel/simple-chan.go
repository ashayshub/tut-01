package main

import "fmt"

func main() {
	c := make(chan bool)
	go func(c chan bool) {
		b := <-c
		fmt.Println(b)
	}(c)
	c <- true
}
