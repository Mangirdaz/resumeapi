package main

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Info(
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
