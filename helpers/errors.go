package helpers

import (
	"log"
)

// HandleError takes an error and message string and logs a fatal error
func HandleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}
