package service

import (
	"log"
)

var URL string = "localhost:5672" 

func FailOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}