package main

import (
	"fmt"
	"solution/solution"
)

func main() {
	fmt.Println("Area:")
	fmt.Println(solution.CalcSquare(4.4, 4))
	fmt.Println(solution.CalcSquare(4.4, 3))
	fmt.Println(solution.CalcSquare(4.4, 0))
}
