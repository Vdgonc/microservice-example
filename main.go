package main

import (
	"log"
	"os"
	"time"
)

func main() {
	var count int64
	NAME := os.Getenv("NAME")

	for {
		log.Printf("Hi Nº %d %s", count+1, NAME)
		time.Sleep(60 * time.Second)
		count++
	}
}
