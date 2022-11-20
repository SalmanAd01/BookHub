package main

import (
	"Bookhub/bin"
	"Bookhub/config"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println("http://localhost:5000/")

	router := bin.CreateServer()
	config.InitRoutes(router)

	const THREESECOND = 3 * time.Second

	server := &http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		ReadHeaderTimeout: THREESECOND,
		Handler:           router,
	}
	log.Fatal(server.ListenAndServe())
}
