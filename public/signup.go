package public

import (
	"Bookhub/models"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	var decoder = schema.NewDecoder()
	var user models.User
	err := decoder.Decode(&user, r.PostForm)
	if err != nil {
		fmt.Println("Error in decoding", err, r.Form.Get("password"))
	}
	salt := []byte(os.Getenv("SALT"))
	hashpwd, err := scrypt.Key([]byte(user.Password), salt, 16384, 8, 1, 32)
	user.Password = hex.EncodeToString(hashpwd)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user.Name, user.Email, user.Password)
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Name,
		Email:    user.Email,
		Password: user.Password,
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
	fmt.Println("--->>> ", r.Host)
	SendMail(user.Email, "Mail Authentication", "http://"+r.Host+"/verifytoken/"+tokenString)
	http.Redirect(w, r, "/signup", http.StatusFound)

}
