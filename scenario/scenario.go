package scenario

import "log"

var count = 0

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

type Camera struct {
	Position  Vector3
	Direction Vector3
}

type Transition interface {
	Camera() Camera
}

type transitionData struct {
	camera Camera
}

func (s transitionData) Camera() Camera {
	return s.camera
}

func Begin() {
	log.Println("Begin")
}

func IsEnd() bool {
	count++
	return count > 2
}

func Progress() Transition {
	log.Println("Progress")
	return transitionData{}
}

func End() {
	log.Println("End")
}

func Skip() {
	count = 100
}
