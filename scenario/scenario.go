package scenario

import "log"

var count = 0

type Vector3 struct {
	x float64
	y float64
	z float64
}

type Camera struct {
	positoin  Vector3
	direction Vector3
}

type Snapshot interface {
	Camera() Camera
}

func Begin() {
	log.Println("Begin")
}

func IsEnd() bool {
	count++
	return count > 2
}

func Progress() Snapshot {
	log.Println("Progress")
	return Snapshot{}
}

func End() {
	log.Println("End")
}
