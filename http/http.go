package goClacks

import "net/http"

func Terrify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Clacks-Overhead", "GNU Terry Pratchett")
		next.ServeHTTP(w, r)
	})
}
