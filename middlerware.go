package main

import "net/http"

type RequestRegister interface {
	handler(next http.Handler) http.Handler
}

var registers = []RequestRegister{
	&requestLog{},
}

func mountParent(next http.Handler) http.Handler {
	n := next
	for _, r := range registers {
		n = r.handler(n)
	}
	return n
}
