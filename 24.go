package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// "Конструктор"
func newPoint(x float64, y float64) Point {
	point := new(Point)
	point.x = x
	point.y = y
	return *point
}

// Функция для подсчета дистанции
func distance(p1 Point, p2 Point) float64 {
	return math.Sqrt(
		math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2),
	)
}

func aaaaaa() {

	a := newPoint(2.7, 10.9)
	b := newPoint(1.7, -10.9)

	fmt.Println(distance(a, b))

}
