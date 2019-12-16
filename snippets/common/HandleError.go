package common

import (
	"fmt"
	"log"
)

// HandleError - Common Utility Function for Package Error Handling
func HandleError(err error) {
	DebugMessage("[common] func HandleError()")
	if err != nil {
		LogMessage("error", fmt.Sprintf("%v", err))
		log.Fatal(err)
	}
}
