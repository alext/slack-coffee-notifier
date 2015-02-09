package main

import (
	"net/http"
	"strings"
)

func NewHandler(slackURL, channel string) http.Handler {
	return &Handler{
		SlackURL: slackURL,
		Channel:  channel,
		client:   &http.Client{},
	}
}

type Handler struct {
	SlackURL string
	Channel  string
	client   *http.Client
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	switch req.Method {
	case "GET":
		h.index(w, req)
	case "POST":
		h.post(w, req)
	default:
		w.Header().Set("Allow", "GET, POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) index(w http.ResponseWriter, req *http.Request, messages ...string) {
	indexTemplate.Execute(w, messages)
}

func (h *Handler) post(w http.ResponseWriter, req *http.Request) {
	message := req.FormValue("message")
	if message == "" {
		h.index(w, req, "Blank message: ignored...")
		return
	}

	url := h.SlackURL + "&channel=%23" + h.Channel

	_, err := h.client.Post(url, "", strings.NewReader(message))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.index(w, req, "Message posted")
}
