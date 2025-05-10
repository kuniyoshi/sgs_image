package scenario

import "log"

var count = 0

func Begin() {
	log.Println("Begin")
}

func IsEnd() bool {
	count++
	return count > 2
}

func Progress() {
	log.Println("Progress")
}

func End() {
	log.Println("End")
}
