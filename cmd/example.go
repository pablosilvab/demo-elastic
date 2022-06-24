package main

import (
	"log"
	"time"

	"github.com/pablosilvab/elastic"
)

type Request struct {
	Request   string
	Method    string
	Timestamp time.Time
}

func main() {
	log.Println("Writing log..")
	err := elastic.Log("test", Request{"/test", "GET", time.Now()})
	if err != nil {
		log.Println("Cannot connect Elastic..")
	}
	log.Println("Hello!")
}
