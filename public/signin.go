package public

import (
	"Bookhub/models"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/scrypt"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SECRET")))

func SigninGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	http.ServeFile(w, r, "./templates/signin.html")
}
func SigninPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	var decoder = schema.NewDecoder()
	var user models.User
	err := decoder.Decode(&user, r.PostForm)
	if err != nil {
		fmt.Println("Error in decoding", err)
	}
	fmt.Println("password --->>", user.Password)
	fmt.Println("username -->>>", user.Name)
	salt := []byte(os.Getenv("SALT"))
	hashpwd, err := scrypt.Key([]byte(user.Password), salt, 16384, 8, 1, 32)
	user.Password = hex.EncodeToString(hashpwd)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("password ", user.Password)
	fmt.Println("email ", user.Email)
	db := models.SetupDB()
	query := "SELECT * FROM userinfo WHERE email = '" + user.Email + "' AND password = '" + user.Password + "'"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Some Error occurred")
		w.Write([]byte("User Already Exists"))
	}
	defer rows.Close()
	if rows.Next() {
		session, _ := Store.Get(r, "auth-session")
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   900,
			HttpOnly: true,
		}
		session.Values["username"] = user.Name
		session.Save(r, w)
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		fmt.Println("No rows")
		w.Write([]byte("Signup unSuccessful"))
	}
	models.CheckErr(err)
}
