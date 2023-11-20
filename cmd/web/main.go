package main

import (
	"github.com/joho/godotenv"
	"log"
)

const portNumber = ":8090"

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	InitServer()
}
