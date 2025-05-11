package main

import (
	"github.com/kuniyoshi/sgs_image/scenario"
)

type vector3 struct {
	x float64
	y float64
	z float64
}

type camera struct {
	positoin  vector3
	directoin vector3
}

type scene struct {
	camera *camera
}

func (scene *scene) sync(snapshot scenario.Snapshot) {

}

type queryType int

const (
	queryTypeUnknown queryType = iota
	queryTypeNext
	queryTypeSkip
)

type player struct {
	query chan queryType
	stop  chan struct{}
}

func (p *player) begin() {
	go func() {
		// TODO: 標準入力を読み込む
		select {
		// TODO: case 標準入力
		// [qQ] なら query に queryTypeSkip を put
		// それ以外なら query に queryTypeNext を put
		case <-p.stop:
			return
		}
	}()
}

func (p *player) end() {
	p.stop <- struct{}{}
}

func main() {
	scene := scene{
		camera: &camera{},
	}

	player := &player{
		query: make(chan queryType),
		stop:  make(chan struct{}),
	}

	scenario.Begin()
	player.begin()

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

		query := <-player.query
		next <- struct{}{}

		if query == queryTypeSkip {
			scenario.Skip()
		}
	}

	player.end()
	scenario.End()
}
