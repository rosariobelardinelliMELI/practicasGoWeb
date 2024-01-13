package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func NewLogger() *Logger {
	return &Logger{}
}

type Logger struct{}

func (l *Logger) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// logic before
		start := time.Now()

		// call next
		next.ServeHTTP(w, r)

		// logic after
		method := r.Method
		end := time.Since(start)
		url := r.URL
		size := r.ContentLength

		fmt.Printf("Verbo utilizado: %s\n", method)
		fmt.Printf("Duración: %d\n", end.Nanoseconds())
		fmt.Printf("URL de la consulta: %s\n", url)
		fmt.Printf("Tamaño en bytes: %d\n", size)
	})
}
