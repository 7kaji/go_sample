package main

import (
	"log"

	"github.com/joho/godotenv"

	db "go_sample/driver"
	"go_sample/server"
)

func main() {
	loadEnv()
	db.Init()
	defer db.Close() // called just before the main function returns
	server.Init()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
