package main

import (
	"github.com/kuniyoshi/sgs_image/scenario"
)

type Vector3 struct {
	x float64
	y float64
	z float64
}

type Camera struct {
	positoin  Vector3
	directoin Vector3
}

type Scene struct {
	camera *Camera
}

func (scene *Scene) sync(snapshot scenario.Snapshot) {

}

func main() {
	scene := Scene{
		camera: &Camera{},
	}

	scenario.Begin()

	for !scenario.IsEnd() {
		snapshot := scenario.Progress()

		go func() {
			defer scene.sync(snapshot)
		}()
	}

	scenario.End()
}
