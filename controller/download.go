package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Download(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookpath := vars["bookpath"]
	fmt.Println("bookpath ", bookpath)
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(bookpath))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, "static/bookinfo/pdf/"+bookpath)
}
