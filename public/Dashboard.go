package public

import (
	"Bookhub/models"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/schema"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/dashboard.html")
}
func DashboardPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DashboardPost")
	err := r.ParseMultipartForm(0)
	if err != nil {
		fmt.Println("Error in parsing form", err)
	}
	var decoder = schema.NewDecoder()
	var book models.Book
	err = decoder.Decode(&book, r.PostForm)
	if err != nil {
		fmt.Println("Error in decoding", err)
	}
	fmt.Println("Book: -->>>", book)
	bookpath, err := SaveFileToDestination(book.Subjectname+"-", book.Semnumber+"-", book.Universityname+"-", r)
	fmt.Println("err ", err, " name ", bookpath)
	imagepath, err := SaveImgToDestination(book.Subjectname+"-", book.Semnumber+"-", book.Universityname+"-", r)
	fmt.Println("err ", err, " name ", imagepath)
	session, err := Store.Get(r, "auth-session")
	if err != nil {
		fmt.Println("Error in getting session", err)
	}
	currentUser := session.Values["userid"].(int)
	session.Save(r, w)
	fmt.Println("currentUser ", reflect.TypeOf(currentUser))
	db := models.SetupDB()
	defer db.Close()
	query := "INSERT INTO bookinfo (bookpath, imgpath, subjectname, bookauthor, semester, branch, universityname, userid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = db.Exec(query, bookpath, imagepath, book.Subjectname, book.Authorname, book.Semnumber, book.Branch, book.Universityname, currentUser)
	if err != nil {
		fmt.Println("Error in inserting bookinfo", err)
	}
	http.Redirect(w, r, "/dashboard", http.StatusFound)

}
