package main

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello")
}
