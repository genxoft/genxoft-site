package middleware

import (
	"github.com/gorilla/handlers"
	"net/http"
)

func CORS(mode string, methods []string) func(http.Handler) http.Handler {
	var o handlers.CORSOption
	if mode == "dev" {
		o = handlers.AllowedOrigins([]string{"*"})
	} else {
		o = handlers.AllowedOrigins([]string{"*"})
	}
	return handlers.CORS(
		o,
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedMethods(methods),
	)
}
