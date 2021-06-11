package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Config struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type H struct{}

func (h *H) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func setDefaults(c *Config) {
	if c.ReadTimeout == 0 {
		c.ReadTimeout = 1
	}
	if c.WriteTimeout == 0 {
		c.WriteTimeout = 1
	}
}

func setConfig(c *Config, options ...func(c *Config)) {
	setDefaults(c)
	for _, option := range options {
		option(c)
	}
}

var setReadTimeout = func(c *Config) {
	c.ReadTimeout = 5
}

var setWriteTimeout = func(c *Config) {
	c.WriteTimeout = 10
}

func NewServer(c *Config) *http.Server {
	var myHandler *H = new(H)
	return &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    c.ReadTimeout * time.Second,
		WriteTimeout:   c.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func main() {
	var c *Config = new(Config)
	setConfig(c)
	s := NewServer(c)
	log.Fatal(s.ListenAndServe())
}
