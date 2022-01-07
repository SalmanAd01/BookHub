package public

import (
	"Bookhub/models"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func VerifyJWT(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("internal error: %v", p)
		}
	}()
	vars := mux.Vars(r)["token"]
	claims := &Claims{}

	fmt.Println("claims -->> ", claims)
	tkn, err := jwt.ParseWithClaims(vars, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("err ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid Token"))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error while parsing token"))
		}
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("err ", err)
		w.Write([]byte("Invalid Token"))
	}
	db := models.SetupDB()
	insertDynStmt := `INSERT INTO userinfo(name,email,password) VALUES($1, $2, $3)`
	_, e := db.Exec(insertDynStmt, claims.Username, claims.Email, claims.Password)
	if e != nil {
		fmt.Println("User already exists")
		w.Write([]byte("User Already Exists"))
	} else {
		fmt.Println("claims -->> ", claims)
		w.Write([]byte("Successful"))
	}
}
