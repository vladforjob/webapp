package main

import (
	"log"
	"net/http"
	"os"

	"bigApp/routes"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	port, exist := os.LookupEnv("PORT")

	if !exist {
		log.Fatal("PORT not set in .env")
	}

	err := http.ListenAndServe(":"+port, routes.Init())

	if err != nil {
		log.Fatal(err)
	}
}
