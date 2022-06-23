package main

import (
	"time"
	"log"

	"github.com/pablosilvab/elastic"
)

type Request struct {
	Request string
	Method  string
	Timestamp time.Time
}

func main() {
	log.Println("Writing log..")
	elastic.Log("test", Request{"/test", "GET", time.Now()})
	log.Println("Ok!")
}
