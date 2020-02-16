package logs

import (
	"log"
	"os"
)

// FailExit err fail exit
func FailExit(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s\n", err, msg)
		os.Exit(1)
	}
}

// FailError fail print error
func FailError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s\n", err, msg)
	}
}
