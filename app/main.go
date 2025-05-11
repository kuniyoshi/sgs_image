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

type QueryType int

const (
	QueryTypeUnknown QueryType = iota
	QueryTypeNext
	QueryTypeSkip
)

type Player interface {
	Query() chan QueryType
}

type Query struct {
}

type mockPlayer struct {
	query chan QueryType
}

func (p *mockPlayer) Query() chan QueryType {
	return p.query
}

func main() {
	scene := Scene{
		camera: &Camera{},
	}

	player := &mockPlayer{query: make(chan QueryType)}

	scenario.Begin()

	for !scenario.IsEnd() {
		snapshot := scenario.Progress()

		next := make(chan struct{})

		go func() {
			defer scene.sync(snapshot)

			for {
				select {
				case <-next:
					return
				}
			}
		}()

		query := <-player.Query()
		next <- struct{}{}

		if query == QueryTypeSkip {
			scenario.Skip()
		}
	}

	scenario.End()
}
