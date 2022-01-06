package main

import (
	"Bookhub/middleware/auth"
	"Bookhub/public"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("http://localhost:5000/")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", auth.IsNotAuth(public.Home)).Methods("GET")
	router.HandleFunc("/signin", auth.IsNotAuth(public.SigninGet)).Methods("GET")
	router.HandleFunc("/signin", auth.IsNotAuth(public.SigninPost)).Methods("POST")
	router.HandleFunc("/signup", auth.IsNotAuth(public.SignupGet)).Methods("GET")
	router.HandleFunc("/signup", auth.IsNotAuth(public.SignupPost)).Methods("POST")
	router.HandleFunc("/verifytoken/{token}", public.VerifyJWT).Methods("GET")
	router.HandleFunc("/dashboard", auth.IsAuth(public.Dashboard)).Methods("GET")
	router.HandleFunc("/dashboard", auth.IsAuth(public.DashboardPost)).Methods("POST")
	router.HandleFunc("/logout", auth.IsAuth(public.Logout)).Methods("GET")
	log.Fatal(http.ListenAndServe(public.GoDotEnvVariable("PORT"), router))

}
