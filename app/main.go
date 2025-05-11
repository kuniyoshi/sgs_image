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
	positoin  vector3
	directoin vector3
}

type scene struct {
	camera *camera
}

func (scene *scene) sync(snapshot scenario.Transition) {

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

		next := make(chan struct{})

		go func() {
			defer scene.sync(transition)

			for {
				select {
				case <-ticker.C:
					scene.tick()
				case <-next:
					return
				}
			}
		}()

		query := <-player.query
		next <- struct{}{}
		close(next)

		if query == queryTypeSkip {
			scenario.Skip()
		}
	}

	player.end()
	scenario.End()
}
