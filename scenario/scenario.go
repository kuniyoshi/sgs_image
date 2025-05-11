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
