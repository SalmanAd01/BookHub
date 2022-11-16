package main

import (
	"Bookhub/bin"
	"Bookhub/config"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	fmt.Println("http://localhost:5000/")
	router := bin.CreateServer()
	config.InitRoutes(router)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
