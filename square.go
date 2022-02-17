package square

import (
	"math"
)

type Sides int8

const (
	SidesTriangle Sides = 3
	SidesSquare   Sides = 4
	SidesCircle   Sides = 0
)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	if sidesNum == SidesTriangle {
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	} else if sidesNum == SidesSquare {
		return sideLen * sideLen
	} else if sidesNum == SidesCircle {
		return math.Pi * sideLen * sideLen
	} else {
		return 0
	}
}
