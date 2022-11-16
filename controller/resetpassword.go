package controller

import (
	"Bookhub/db"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ResetPasswordGet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	email := r.Form.Get("emailforget")
	fmt.Println("email ", email)
	db := db.Connect()
	query := "SELECT id FROM userinfo WHERE email = '" + email + "'"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Some Error occurred")
	}
	defer rows.Close()
	if rows.Next() {
		fmt.Println("rows ", rows)
		// w.Write([]byte("Check your email for reset password link"))
	} else {
		w.Write([]byte("Email not found"))
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	fmt.Println("claims ", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT")))
	if err != nil {
		fmt.Println("err ", err)
		return
	}
	fmt.Println("token ", tokenString)
	SendMail(email, "Reset Password", "http://"+r.Host+"/forgotpassword/"+tokenString)
	http.Redirect(w, r, "/signin", http.StatusFound)

}
