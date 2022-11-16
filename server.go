package main

import (
	"Bookhub/controller"
	"Bookhub/middleware/auth"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	fmt.Println("http://localhost:5000/")
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	router.HandleFunc("/", (controller.Home)).Methods("GET")
	router.HandleFunc("/signin", auth.IsNotAuth(controller.SigninGet)).Methods("GET")
	router.HandleFunc("/signin", auth.IsNotAuth(controller.SigninPost)).Methods("POST")
	router.HandleFunc("/signup", auth.IsNotAuth(controller.SignupGet)).Methods("GET")
	router.HandleFunc("/signup", auth.IsNotAuth(controller.SignupPost)).Methods("POST")
	router.HandleFunc("/verifytoken/{token}", controller.VerifyJWT).Methods("GET")
	router.HandleFunc("/forgotpassword/{token}", auth.IsNotAuth(controller.ForgotPasswordGet)).Methods("GET")
	router.HandleFunc("/forgotpassword/{token}", auth.IsNotAuth(controller.ForgotPasswordPost)).Methods("POST")
	router.HandleFunc("/resetpassword", auth.IsNotAuth(controller.ResetPasswordGet)).Methods("POST")
	router.HandleFunc("/dashboard", auth.IsAuth(controller.Dashboard)).Methods("GET")
	router.HandleFunc("/dashboard", auth.IsAuth(controller.DashboardPost)).Methods("POST")
	router.HandleFunc("/logout", auth.IsAuth(controller.Logout)).Methods("GET")
	router.HandleFunc("/download/{bookpath}", controller.Download).Methods("GET")
	router.HandleFunc("/delete/{bookpath}", (controller.Delete)).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
