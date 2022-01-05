package public

import "net/http"

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "auth-session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/signin", 302)
}
