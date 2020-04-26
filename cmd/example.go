package main

import (
	"log"

	"github.com/pablosilvab/elastic-lib"
)

type Request struct {
	Request string
	Method  string
}

func main() {
	log.Println("Writing log..")
	elastic.Log("test", Request{"/test", "POST"})
	log.Println("Ok!")
}
