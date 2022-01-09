package public

import (
	"Bookhub/models"
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	db := models.SetupDB()
	query := "SELECT bookpath,imgpath,subjectname,semester,universityname,branch FROM bookinfo"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error in getting bookinfo", err)
	}
	bookinfo := []models.Book{}
	defer rows.Close()
	var book models.Book
	for rows.Next() {
		fmt.Println("book ")

		err = rows.Scan(&book.Bookfile, &book.Bannerimage, &book.Subjectname, &book.Semnumber, &book.Universityname, &book.Branch)
		if err != nil {
			fmt.Println("Error in scanning bookinfo", err)
		}
		bookinfo = append(bookinfo, book)

		fmt.Println("bookinfo ", bookinfo)
	}
	t, err := template.ParseFiles("./views/index.html")
	if err != nil {
		fmt.Println("Error in parsing home.html", err)
	}
	t.Execute(w, bookinfo)
	// http.ServeFile(w, r, "./views/index.html")
}

// 	t, err := template.ParseFiles("views/index.html")
// 	if err != nil {
// 		fmt.Println("Error in parsing template", err)
// 	}
// 	t.Execute(w, []int{1, 2, 3})
// 	// fmt.Println("Home")
// }
