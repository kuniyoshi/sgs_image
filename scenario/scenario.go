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

type TransitionType int

const (
	TransitionTypeUnknown TransitionType = iota
)

type Transition struct {
	Camera Camera
}

func Begin() {
	log.Println("Begin")
}

func IsEnd() bool {
	count++
	return count > 2
}

func Progress() []Transition {
	log.Println("Progress")

	// 最終的なカメラの状態を表す複数のTransitionを返す
	transitions := make([]Transition, 0)

	switch count % 2 {
	case 0:
		// 右側の位置
		transitions = append(transitions, Transition{
			Camera: Camera{
				Position: Vector3{
					X: 1,
				},
				Direction: Vector3{
					Z: -1,
				},
			},
		})
	case 1:
		// 左側の位置
		transitions = append(transitions, Transition{
			Camera: Camera{
				Position: Vector3{
					X: -1,
				},
				Direction: Vector3{
					Z: -1,
				},
			},
		})
	}

	return transitions
}

func End() {
	log.Println("End")
}

func Skip() {
	count = 100
}
