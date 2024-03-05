package utils

import (
	"encoding/json"
	"log"
)

func LogError(err interface{}) {
	if err != nil {
		msg, _ := json.Marshal(err)
		log.Printf("ERROR: %s\n", string(msg))
	}
}
