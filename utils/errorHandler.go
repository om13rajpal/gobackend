package utils

import "log"

func HandleError(err error, errorMsg string) {
	if err != nil {
		log.Fatal(errorMsg, "\n", err)
	}
}
