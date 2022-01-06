package public

import (
	"Bookhub/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/schema"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/dashboard.html")
}
func DashboardPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DashboardPost")
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error in parsing form", err)
	}
	_, file, err := r.FormFile("bookfile")
	if err != nil {
		fmt.Println("Error in getting file", err)
		return
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println("Error in opening file", err)
		return
	}
	defer src.Close()

	dst, err := os.Create(filepath.Join("C:/Users/salman/OneDrive/Desktop/BookHub/static/pdf", filepath.Base(file.Filename))) // dir is directory where you want to save file.
	if err != nil {
		fmt.Println("Error in creating file", err)
		return
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("Error in copying file", err)
		return
	}
	if err != nil {
		fmt.Println("Error in getting bookfile", err)
	}
	// fmt.Println("bookfile", bookfile, "bookfilename", bookfilename)
	var decoder = schema.NewDecoder()
	var book models.Book
	err = decoder.Decode(&book, r.PostForm)
	if err != nil {
		fmt.Println("Error in decoding", err)
	}
	fmt.Println("Book: ", book)
	http.Redirect(w, r, "/dashboard", http.StatusFound)

}
