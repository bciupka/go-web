package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello from middleware")
		next.ServeHTTP(w, r)
	})
}

// NoSurf creates CSRG protecion for all POSTs to a server
func NoSurf(next http.Handler) http.Handler {
	csrfToken := nosurf.New(next)
	csrfToken.SetBaseCookie(http.Cookie{
		Path:     "/",
		Secure:   app.InProduction,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfToken
}

// SessionLoad loads a session every time the page is being reloaded
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
