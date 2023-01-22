package main

import (
	"github.com/corentings/kafejo/app/http"
	"github.com/corentings/kafejo/infrastructures"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Printf("Hello, world.")
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = infrastructures.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = infrastructures.CloseDB()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Create the app
	app := http.New()
	// Listen to port 1812
	log.Panic(app.Listen(":1819"))

}
