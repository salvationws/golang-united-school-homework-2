package solution

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sidesInt int

func CalcSquare(sideLen float64, sidesNum sidesInt) float64 {
	switch sidesNum {
	case 0:
		fmt.Print("Circle with Radius: ")
		size := math.Pi * sideLen * 2.0
		fmt.Print(size)
		fmt.Print(" and Diameter: ")
		fmt.Print(sideLen * 2.0)
	case 3:
		fmt.Print("Equilateral Triangle")
	case 4:
		fmt.Print("Square")
	default:
		return 0
	}
	return float64(sidesNum)
}

func main() {
	fmt.Println("Your figure is:")
	CalcSquare(5.0, 0)
}
