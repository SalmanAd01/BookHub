package controller

import (
	"Bookhub/db"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookpath := vars["bookpath"]
	name := r.URL.Query().Get("name")
	password := r.URL.Query().Get("password")
	fileType := r.URL.Query().Get("type")
	isDB := r.URL.Query().Get("db")

	if name != os.Getenv("ADMIN_NAME") || password != os.Getenv("ADMIN_PASSWORD") || (fileType != "pdf" && fileType != "img") {
		fmt.Fprintf(w, "Invalid credentials")
		return
	}

	file, err := os.Getwd()

	if err != nil {
		fmt.Fprintf(w, "Error in getting file")
		return
	}

	var PATH = file + "/static/bookinfo/" + fileType + "/" + bookpath

	err = os.Remove(PATH)

	if err != nil {
		fmt.Fprintf(w, "Error deleting file "+err.Error())
		return
	}

	if isDB == "true" {
		dbs := db.Connect()

		defer dbs.Close()

		query := "DELETE FROM bookinfo WHERE bookpath = $1"
		_, err = dbs.Exec(query, bookpath)

		if err != nil {
			fmt.Fprintf(w, "Error deleting from database")
			return
		}
	}

	fmt.Println("error", err)
	fmt.Fprintf(w, "File deleted successfully")
}
