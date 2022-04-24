package tools

import (
	"log"
	"os"
)

func Err(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
