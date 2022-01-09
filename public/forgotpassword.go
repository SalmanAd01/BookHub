package public

import (
	"Bookhub/models"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/scrypt"
)

func ForgotPasswordGet(w http.ResponseWriter, r *http.Request) {
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
	tmpToken := struct {
		Name string
	}{
		Name: vars,
	}
	t, err := template.ParseFiles("templates/forgotpassword.html")
	if err != nil {
		fmt.Println("err ", err)
	}
	fmt.Println("claims -->> ", claims.Email, t)
	t.Execute(w, tmpToken)
}
func ForgotPasswordPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)["token"]
	claims := &Claims{}

	fmt.Println("claims -->> ", claims)
	tkn, err := jwt.ParseWithClaims(vars, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("err ", err)
			w.WriteHeader(http.StatusFound)
			w.Write([]byte("Invalid Token"))
		} else {
			w.WriteHeader(http.StatusFound)
			w.Write([]byte("Error while parsing token"))
		}
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusFound)
		fmt.Println("err ", err)
		w.Write([]byte("Invalid Token"))
	}
	fmt.Println("claims -->> ", claims.Email)
	err = r.ParseForm()
	if err != nil {
		fmt.Println("err ", err)
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("Error while parsing form"))
	}
	password1 := r.Form.Get("password1")
	password2 := r.Form.Get("password2")
	fmt.Println("password ", password1, " ", password2)
	if password1 != password2 {
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("Passwords do not match"))
	} else {
		fmt.Println("claims -->> ", claims.Email)
		salt := []byte(os.Getenv("SALT"))
		hashpwd, err := scrypt.Key([]byte(password1), salt, 16384, 8, 1, 32)
		password1 = hex.EncodeToString(hashpwd)
		if err != nil {
			log.Println(err)
		}
		db := models.SetupDB()
		defer db.Close()
		query := "UPDATE userinfo SET password = $1 WHERE email = $2"
		_, err = db.Exec(query, password1, claims.Email)
		if err != nil {
			log.Println(err)
		}
		// w.Write([]byte("Password changed successfully"))
		http.Redirect(w, r, "/signin", http.StatusFound)
	}
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Password updated successfully"))

	// password:=r.FormValue("password")
}
