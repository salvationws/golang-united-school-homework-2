package solution

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sidesInt int

var SidesTriangle sidesInt = 3
var SidesSquare sidesInt = 4
var SidesCircle sidesInt = 0

func CalcSquare(sideLen float64, sidesNum sidesInt) float64 {
	if sidesNum == 4 {
		SidesSquare := (sideLen * sideLen)
		return SidesSquare
	} else if sidesNum == 3 {
		SidesTriangle := ((sideLen * sideLen) * math.Sqrt(3)) / 4
		return SidesTriangle
	} else if sidesNum == 0 {
		SidesCircle := math.Pi * (sideLen * sideLen)
		return SidesCircle
	}

	return 0
}
