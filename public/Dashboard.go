package public

import "net/http"

func Dashboard(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/dashboard.html")
}
