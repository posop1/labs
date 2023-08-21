package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getCircleArea(2))
}

func getCircleRadius(radius float64) float64 {
	pi := 3.14

	result := 2 * pi * radius

	return result
}

func getCircleArea(radius float64) float64 {
	pi := 3.14

	result := pi * math.Pow(radius, 2)

	return result
}
