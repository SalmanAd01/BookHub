package main

import (
	public "Bookhub/public"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("http://localhost:5000/")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", public.Home).Methods("GET")
	router.HandleFunc("/signin", public.SigninGet).Methods("GET")
	router.HandleFunc("/signin", public.SigninPost).Methods("POST")
	router.HandleFunc("/signup", public.SignupGet).Methods("GET")
	router.HandleFunc("/signup", public.SignupPost).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))

}
