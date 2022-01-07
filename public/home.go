package public

import (
	"Bookhub/models"
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type tempStrcut struct {
	Name     string
	Image    string
	Location string
}

func Home(w http.ResponseWriter, r *http.Request) {
	db := models.SetupDB()
	query := "SELECT bookpath,imgpath,subjectname FROM bookinfo"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error in getting bookinfo", err)
	}
	bookinfo := []tempStrcut{}
	defer rows.Close()
	var book tempStrcut
	for rows.Next() {
		fmt.Println("book ")

		err = rows.Scan(&book.Location, &book.Image, &book.Name)
		if err != nil {
			fmt.Println("Error in scanning bookinfo", err)
		}
		book.Image = strings.Replace(book.Image, "\\", "/", -1)
		book.Location = strings.Replace(book.Location, "\\", "/", -1)
		// book.Location = book.Location[60:]
		// book.Image = book.Image[40:]
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
