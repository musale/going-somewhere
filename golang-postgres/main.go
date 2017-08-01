package main

import (
	"net/http"

	"github.com/musale/going-somewhere/golang-postgres/lib"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "peng",
	}).Info("A peng appears")

	http.HandleFunc("/", lib.MessagesPage)
	http.ListenAndServe(":5505", nil)
}
