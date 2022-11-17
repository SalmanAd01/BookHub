package controller

import (
	"Bookhub/models"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/schema"
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
	fmt.Println("POST")

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	var (
		decoder = schema.NewDecoder()
		user    models.User
	)

	err := decoder.Decode(&user, r.PostForm)

	if err != nil {
		fmt.Println("Error in decoding", err, r.Form.Get("password"))
	}

	salt := []byte(os.Getenv("SALT"))

	const (
		MEMORYCOST = 16384
		THREADS    = 8
		KEYLENGTH  = 32
	)

	hashpwd, err := scrypt.Key([]byte(user.Password), salt, MEMORYCOST, THREADS, 1, KEYLENGTH)
	user.Password = hex.EncodeToString(hashpwd)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(user.Name, user.Email, user.Password)

	const FIVEMINUTES = 5 * time.Minute
	expirationTime := time.Now().Add(FIVEMINUTES)
	claims := jwt.MapClaims{
		"Username": user.Name,
		"Email":    user.Email,
		"Password": user.Password,
		"exp":      expirationTime.Unix(),
	}
	fmt.Println("claims ", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT")))

	if err != nil {
		fmt.Println("err ", err)
		return
	}

	fmt.Println("token ", tokenString)
	fmt.Println("--->>> ", r.Host)
	SendMail(user.Email, "Mail Authentication", "http://"+r.Host+"/verifytoken/"+tokenString)
	http.Redirect(w, r, "/signup", http.StatusFound)
}
