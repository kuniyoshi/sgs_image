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

func Progress() Transition {
	log.Println("Progress")
	switch count % 2 {
	case 0:
		return Transition{
			Camera: Camera{
				Position: Vector3{
					X: 1,
				},
				Direction: Vector3{
					Z: -1,
				},
			},
		}
	case 1:
		return Transition{
			Camera: Camera{
				Position: Vector3{
					X: -1,
				},
				Direction: Vector3{
					Z: -1,
				},
			},
		}
	}
	panic("unreachable")
}

func End() {
	log.Println("End")
}

func Skip() {
	count = 100
}
