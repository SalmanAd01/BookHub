package public

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	http.ServeFile(w, r, "./templates/index.html")
}
