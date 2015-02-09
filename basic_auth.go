package main

import "net/http"

type BasicAuth struct {
	username string
	password string
	next     http.Handler
}

func NewBasicAuth(username, password string, next http.Handler) http.Handler {
	return &BasicAuth{username, password, next}
}

func (ba *BasicAuth) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	u, p, ok := req.BasicAuth()
	if !ok || u != ba.username || p != ba.password {
		w.Header().Set("WWW-Authenticate", "Basic realm=\"Authorization Required\"")
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return
	}
	ba.next.ServeHTTP(w, req)
}
