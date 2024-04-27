package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf is a middleware function that adds CSRF protection to all routes.
// It wraps the provided http.Handler and returns a new http.Handler that includes CSRF protection.
// The CSRF protection is implemented using the nosurf package.
// The CSRF cookie is set with HttpOnly flag, Path set to "/", Secure flag based on app.InProduction, and SameSite set to http.SameSiteLaxMode.
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

// SessionLoad is a middleware function that loads and saves session data for each request.
// It takes an http.Handler as input and returns an http.Handler.
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
