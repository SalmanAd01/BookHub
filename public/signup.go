package public

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/scrypt"
)

type Claims struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func SignupGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	http.ServeFile(w, r, "./templates/signup.html")

}

func SignupPost(w http.ResponseWriter, r *http.Request) {
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
	salt := []byte(GoDotEnvVariable("SALT"))
	hashpwd, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)
	password = hex.EncodeToString(hashpwd)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(name, email, password)
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
	http.Redirect(w, r, "/signup", http.StatusFound)

}
