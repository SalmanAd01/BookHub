package public

import (
	"Bookhub/models"
	"fmt"
	"net/http"

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
	var decoder = schema.NewDecoder()
	var book models.Book
	err = decoder.Decode(&book, r.PostForm)
	if err != nil {
		fmt.Println("Error in decoding", err)
	}
	fmt.Println("Book: ", book)
	http.Redirect(w, r, "/dashboard", http.StatusFound)

}
