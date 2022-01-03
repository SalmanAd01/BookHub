package public

import (
	"Bookhub/models"
	"fmt"
	"net/http"
)

func SignupGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	http.ServeFile(w, r, "./templates/signup.html")
}
func SignupPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println("email ", email)
	fmt.Println("password ", password)
	db := models.SetupDB()
	query := "SELECT * FROM userinfo WHERE email = '" + email + "' AND password = '" + password + "'"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Some Error occurred")
		w.Write([]byte("User Already Exists"))
	}
	defer rows.Close()
	if rows.Next() {
		http.Redirect(w, r, "/signup", http.StatusFound)
	} else {
		fmt.Println("No rows")
		w.Write([]byte("Signup unSuccessful"))
	}
	models.CheckErr(err)
}
