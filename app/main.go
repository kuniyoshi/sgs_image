package main

import (
	"log"

	"github.com/kuniyoshi/sgs_image/scenario"
)

func main() {
	scenario.Begin()

	for !scenario.IsEnd() {
		scenario.Progress()
	}

	log.Println("DONE")
}
