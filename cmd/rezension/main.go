package main

import (
	"flag"
	"log"
)

var (
	timestamps bool
)

func main() {
	flag.BoolVar(&timestamps, "timestamps", true, "print timestamps in output")
	flag.Parse()

	if !timestamps {
		log.SetFlags(0)
	}
	log.Println("Startup rezenssion monolith")
}
