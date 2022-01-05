package auth

import (
	"Bookhub/public"
	"fmt"
	"net/http"
)

func IsAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := public.Store.Get(r, "auth-session")
		if username, ok := session.Values["username"].(string); ok {
			fmt.Println("username:-->>> ", username)
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/signin", http.StatusUnauthorized)
		}
	})
}
func IsNotAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := public.Store.Get(r, "auth-session")
		if username, ok := session.Values["username"].(string); ok {
			fmt.Println("username:-->>> ", username)
			http.Redirect(w, r, "/dashboard", http.StatusUnauthorized)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
