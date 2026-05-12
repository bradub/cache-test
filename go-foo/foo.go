// Package foo provides a small HTTP router used to exercise CI caching.
package foo

import (
	"html"
	"net/http"
	"os"

	"github.com/aarondl/opt/null"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

// Greet returns a greeting for the optional name.
func Greet(name null.Val[string]) string {
	value, ok := name.Get()
	if !ok {
		return "hello, stranger"
	}

	return "hello, " + value
}

// Logger returns a zerolog logger writing to stderr.
func Logger() zerolog.Logger {
	return zerolog.New(os.Stderr).With().Timestamp().Logger()
}

// NewRouter returns a chi router with a single /hello endpoint.
func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/hello", helloHandler)

	return router
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	var name null.Val[string]

	if query := request.URL.Query().Get("name"); query != "" {
		name = null.From(query)
	}

	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = writer.Write([]byte(html.EscapeString(Greet(name))))
}
