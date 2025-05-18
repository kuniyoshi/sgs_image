package main

import (
	"bufio"
	"fmt"
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

func (v vector3) String() string {
	return fmt.Sprintf("{x: %.2f, y: %.2f, z: %.2f}", v.x, v.y, v.z)
}

type camera struct {
	position  vector3
	direction vector3
}

func (c camera) String() string {
	return fmt.Sprintf("{position: %s, direction: %s}", c.position.String(), c.direction.String())
}

type scene struct {
	camera *camera
}

func (s scene) String() string {
	return fmt.Sprintf("{camera: %s}", s.camera.String())
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
	log.Printf("%+v", scene)
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
		transitions := scenario.Progress()

		// 各Transitionを目標位置として処理
		for _, targetTransition := range transitions {
			// 現在のカメラ位置を保存
			startCamera := camera{
				position:  scene.camera.position,
				direction: scene.camera.direction,
			}

			// 目標位置
			targetCamera := camera{
				position: vector3{
					x: targetTransition.Camera.Position.X,
					y: targetTransition.Camera.Position.Y,
					z: targetTransition.Camera.Position.Z,
				},
				direction: vector3{
					x: targetTransition.Camera.Direction.X,
					y: targetTransition.Camera.Direction.Y,
					z: targetTransition.Camera.Direction.Z,
				},
			}

			// アニメーションの総ステップ数
			const totalSteps = 10
			currentStep := 0

			// アニメーションループ
		animationLoop:
			for currentStep <= totalSteps {
				select {
				case <-ticker.C:
					// 現在のステップに基づいて位置を補間
					t := float64(currentStep) / float64(totalSteps)

					// 線形補間
					scene.camera.position = vector3{
						x: startCamera.position.x + t*(targetCamera.position.x-startCamera.position.x),
						y: startCamera.position.y + t*(targetCamera.position.y-startCamera.position.y),
						z: startCamera.position.z + t*(targetCamera.position.z-startCamera.position.z),
					}

					scene.camera.direction = vector3{
						x: startCamera.direction.x + t*(targetCamera.direction.x-startCamera.direction.x),
						y: startCamera.direction.y + t*(targetCamera.direction.y-startCamera.direction.y),
						z: startCamera.direction.z + t*(targetCamera.direction.z-startCamera.direction.z),
					}

					scene.tick()
					currentStep++

					// アニメーション完了
					if currentStep > totalSteps {
						break animationLoop
					}
				case query := <-player.query:
					if query == queryTypeSkip {
						scenario.Skip()
						// スキップの場合は即座に目標位置に設定
						scene.camera.position = targetCamera.position
						scene.camera.direction = targetCamera.direction
						scene.tick()
						break animationLoop
					}
				}
			}

			// ユーザー入力待ち
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
		}
	}

	player.end()
	scenario.End()
}
