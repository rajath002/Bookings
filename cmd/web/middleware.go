package main

import (
	"net/http"
	"strings"

	"github.com/justinas/nosurf"
	"github.com/rajath002/bookings/internal/helpers"
)

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

// Loads and saves the session in every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func LogAccessingUrl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/static/") {
			next.ServeHTTP(w, r)
			return
		}
		app.InfoLog.Println(helpers.GetFullURL(r))
		next.ServeHTTP(w, r)
	})
}
