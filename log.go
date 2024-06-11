package main

import (
	"fmt"
	"net/http"
)

type requestLog struct{}

func (*requestLog) handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := redDA.dyeing(r.RemoteAddr)
		method := greenDA.dyeing(r.Method)
		path := blueDA.dyeing(r.URL.Path)
		ua := yellowDA.dyeing(r.UserAgent())
		fmt.Printf("Received request: Remote Address: %s, Method: %s, Path: %s, User-Agent: %s\n",
			addr, method, path, ua)
		next.ServeHTTP(w, r)
	})
}
