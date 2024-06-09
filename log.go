package main

import (
	"fmt"
	"net/http"
)

func logMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: Remote Address: %s, Method: %s, Path: %s, User-Agent: %s, Headers: %+v\n",
			r.RemoteAddr, r.Method, r.URL.Path, r.UserAgent(), r.Header)
		next.ServeHTTP(w, r)
	})
}
