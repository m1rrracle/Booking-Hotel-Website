package main

import (
	"github.com/justinas/nosurf"
	"net/http"
	"pdev-finmart/internal/helpers"
)

// NoSurf is the csrf protection middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves session data for current request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func Auth(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuthenticate(r){
			session.Put(r.Context(), "error", "Log in first!")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		}
		next.ServeHTTP(w, r)
	})
}