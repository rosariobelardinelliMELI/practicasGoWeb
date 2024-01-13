package middleware

import (
	"net/http"
)

func NewAutenticator(token string) *Authenticator {
	return &Authenticator{
		token: token,
	}
}

type Authenticator struct {
	token string
}

func (a *Authenticator) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// logic before
		tokenHeader := r.Header.Get("token")
		// println("Token header:", tokenHeader)
		// println("Token:", h.token)
		if tokenHeader == "" {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing token"))
			return
		}
		if tokenHeader != a.token {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Wrong token"))
			return
		}

		// call next
		next.ServeHTTP(w, r)

		// logic after
		// ...
	})
}
