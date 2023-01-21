package main

import (
	"github.com/memnix/memnix-rest/data/infrastructures"
	"log"
)

func main() {
	log.Printf("Hello, world.")

	err := infrastructures.Connect()
	if err != nil {
		return
	}

	defer infrastructures.Disconnect()

}
