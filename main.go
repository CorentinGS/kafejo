package main

import (
	"github.com/corentings/kafejo/infrastructures"
	"log"
)

func main() {
	log.Printf("Hello, world.")

	err := infrastructures.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := infrastructures.CloseDB()
		if err != nil {
			log.Fatal(err)
		}
	}()

}
