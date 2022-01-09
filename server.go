package main

import (
	"Bookhub/middleware/auth"
	"Bookhub/public"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	fmt.Println("http://localhost:5000/")
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", auth.IsNotAuth(public.Home)).Methods("GET")
	router.HandleFunc("/signin", auth.IsNotAuth(public.SigninGet)).Methods("GET")
	router.HandleFunc("/signin", auth.IsNotAuth(public.SigninPost)).Methods("POST")
	router.HandleFunc("/signup", auth.IsNotAuth(public.SignupGet)).Methods("GET")
	router.HandleFunc("/signup", auth.IsNotAuth(public.SignupPost)).Methods("POST")
	router.HandleFunc("/verifytoken/{token}", public.VerifyJWT).Methods("GET")
	router.HandleFunc("/forgotpassword/{token}", auth.IsNotAuth(public.ForgotPasswordGet)).Methods("GET")
	router.HandleFunc("/forgotpassword/{token}", auth.IsNotAuth(public.ForgotPasswordPost)).Methods("POST")
	router.HandleFunc("/resetpassword", auth.IsNotAuth(public.ResetPasswordGet)).Methods("POST")
	router.HandleFunc("/dashboard", auth.IsAuth(public.Dashboard)).Methods("GET")
	router.HandleFunc("/dashboard", auth.IsAuth(public.DashboardPost)).Methods("POST")
	router.HandleFunc("/logout", auth.IsAuth(public.Logout)).Methods("GET")
	router.HandleFunc("/download/{bookpath}", public.Download).Methods("GET")
	router.HandleFunc("/delete/{bookpath}", (public.Delete)).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
