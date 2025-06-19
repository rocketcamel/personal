package middleware

import (
	"context"
	"net/http"
)

func ThemeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var theme string
		cookie, err := r.Cookie("theme")
		if err == nil && cookie.Value == "dark" {
			theme = "dark"
		}

		ctx := context.WithValue(r.Context(), "theme", theme)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
