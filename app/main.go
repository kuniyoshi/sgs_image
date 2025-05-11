package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/kuniyoshi/sgs_image/scenario"
)

type vector3 struct {
	x float64
	y float64
	z float64
}

type camera struct {
	position  vector3
	direction vector3
}

type scene struct {
	camera *camera
}

func (scene *scene) sync(transition scenario.Transition) {
	scene.camera.position = vector3{
		x: transition.Camera.Position.X,
		y: transition.Camera.Position.Y,
		z: transition.Camera.Position.Z,
	}
	scene.camera.direction = vector3{
		x: transition.Camera.Direction.X,
		y: transition.Camera.Direction.Y,
		z: transition.Camera.Direction.Z,
	}
}

func (scene *scene) tick() {
	log.Println("tick")
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
	inputChan := make(chan string)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			select {
			case <-p.stop:
				close(inputChan)
				return
			default:
				line, err := reader.ReadString('\n')
				if err != nil {
					close(inputChan)
					return
				}
				inputChan <- line
			}
		}
	}()

	go func() {
		for {
			select {
			case <-p.stop:
				return
			case line, ok := <-inputChan:
				if !ok {
					return
				}
				switch strings.ToLower(strings.TrimSpace(line)) {
				case "q":
					p.query <- queryTypeSkip
				default:
					p.query <- queryTypeNext
				}
			}
		}
	}()
}

func (p *player) end() {
	p.stop <- struct{}{}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("--- ")

	scene := scene{
		camera: &camera{},
	}

	player := &player{
		query: make(chan queryType),
		stop:  make(chan struct{}),
	}

	scenario.Begin()
	player.begin()
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for !scenario.IsEnd() {
		transition := scenario.Progress()

	queryLoop:
		for {
			select {
			case <-ticker.C:
				scene.tick()
			case query := <-player.query:
				if query == queryTypeSkip {
					scenario.Skip()
				}
				break queryLoop
			}
		}

		scene.sync(transition)
	}

	player.end()
	scenario.End()
}
