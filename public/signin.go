package public

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func SigninGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	http.ServeFile(w, r, "./templates/signin.html")

}

func SigninPost(w http.ResponseWriter, r *http.Request) {
	// defer func() {
	// 	if p := recover(); p != nil {
	// 		fmt.Printf("internal error: %v", p)
	// 	}
	// }()
	fmt.Println("POST")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	name, email, password := r.FormValue("name"), r.FormValue("email"), r.FormValue("password")
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: name,
		Email:    email,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	fmt.Println("claims ", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(GoDotEnvVariable("JWT")))
	if err != nil {
		fmt.Println("err ", err)
		return
	}
	fmt.Println("token ", tokenString)
	fmt.Println("--->>> ", r.Host)
	SendMail(email, "Mail Authentication", "http://"+r.Host+"/verifytoken/"+tokenString)
	http.Redirect(w, r, "/signin", http.StatusFound)

}
