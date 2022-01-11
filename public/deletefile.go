package public

import (
	"Bookhub/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookpath := vars["bookpath"]
	name := r.URL.Query().Get("name")
	password := r.URL.Query().Get("password")
	fileType := r.URL.Query().Get("type")
	isDB := r.URL.Query().Get("db")
	fmt.Println(bookpath, " ", name, " ", password)
	if name != os.Getenv("ADMIN_NAME") || password != os.Getenv("ADMIN_PASSWORD") || (fileType != "pdf" && fileType != "img") {
		fmt.Fprintf(w, "Invalid credentials")
		return
	}
	file, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(w, "Error in getting file")
		return
	}
	err = os.Remove(filepath.Join(file + "/static/bookinfo/" + fileType + "/" + bookpath))
	if err != nil {
		fmt.Fprintf(w, "Error deleting file "+err.Error())
		return
	}
	if isDB == "true" {
		db := models.SetupDB()
		defer db.Close()
		query := "DELETE FROM bookinfo WHERE bookpath = $1"
		_, err = db.Exec(query, bookpath)
		if err != nil {
			fmt.Fprintf(w, "Error deleting from database")
			return
		}
	}

	fmt.Println("error", err)
	fmt.Fprintf(w, "File deleted successfully")
}
