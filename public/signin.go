package public

import (
	"Bookhub/models"
	"fmt"
	"net/http"
)

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
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println("Name ", name)
	fmt.Println("email ", email)
	fmt.Println("Password ", password)
	db := models.SetupDB()
	insertDynStmt := `INSERT INTO userinfo(name,email,password) VALUES($1, $2, $3)`
	_, e := db.Exec(insertDynStmt, name, email, password)
	if e != nil {
		fmt.Println("User already exists")
	}
	http.Redirect(w, r, "/signin", http.StatusFound)

}
