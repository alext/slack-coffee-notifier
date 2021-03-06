package main

import (
	"log"
	"net/http"
	"os"
)

var (
	port         = getEnvDefault("PORT", "8080")
	username     = os.Getenv("HTTP_USERNAME")
	password     = os.Getenv("HTTP_PASSWORD")
	slackURL     = os.Getenv("SLACK_URL")
	slackChannel = os.Getenv("SLACK_CHANNEL")
)

func main() {
	if slackURL == "" || slackChannel == "" {
		log.Fatal("Must specify SLACK_URL and SLACK_CHANNEL env vars")
	}

	rootHandler := NewHandler(slackURL, slackChannel)

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
