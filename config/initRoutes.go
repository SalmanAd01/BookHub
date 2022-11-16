package config

import (
	"Bookhub/controller"
	"Bookhub/middleware/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
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
}
