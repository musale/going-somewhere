package lib

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// MessagesPage - handle some message
func MessagesPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.WithField("request", "received")
	w.Header().Set("Content-Type", "application/json")
	return
}
