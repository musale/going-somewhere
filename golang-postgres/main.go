package main

import(
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.WithFields(log.Fields{
    "animal": "peng",
  }).Info("A peng appears")
}
