package main

import "fmt"

func main() {

	fmt.Println(getCircleRadius())

}

func getCircleRadius() float64 {
	pi := 3.14
	var radius float64

	fmt.Println("Введите радиус окружности. ")

	fmt.Scan(&radius)

	result := (radius * radius) * pi

	return result
}
