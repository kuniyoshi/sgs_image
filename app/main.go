package main

import (
	"github.com/kuniyoshi/sgs_image/scenario"
)

func main() {
	scenario.Begin()

	for !scenario.IsEnd() {
		scenario.Progress()
	}

	scenario.End()
}
