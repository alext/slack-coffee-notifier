package main

import (
	"log"
	"net/http"
	"os"
)

var (
	port     = getEnvDefault("PORT", "8080")
	username = os.Getenv("HTTP_USERNAME")
	password = os.Getenv("HTTP_PASSWORD")
)

func main() {
	var rootHandler http.Handler
	rootHandler = &Handler{}

	if username != "" && password != "" {
		log.Println("Adding basic auth")
		rootHandler = NewBasicAuth(username, password, rootHandler)
	}

	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, rootHandler))
}

func getEnvDefault(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}

	return val
}
