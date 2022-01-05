package public

import (
	"Bookhub/models"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/scrypt"
)

func SigninGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	http.ServeFile(w, r, "./templates/signin.html")
}
func SigninPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println("password ", password)
	salt := []byte(GoDotEnvVariable("SALT"))
	hashpwd, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)
	password = hex.EncodeToString(hashpwd)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("password ", password)
	fmt.Println("email ", email)
	db := models.SetupDB()
	query := "SELECT * FROM userinfo WHERE email = '" + email + "' AND password = '" + password + "'"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Some Error occurred")
		w.Write([]byte("User Already Exists"))
	}
	defer rows.Close()
	if rows.Next() {
		http.Redirect(w, r, "/signin", http.StatusFound)
	} else {
		fmt.Println("No rows")
		w.Write([]byte("Signup unSuccessful"))
	}
	models.CheckErr(err)
}
