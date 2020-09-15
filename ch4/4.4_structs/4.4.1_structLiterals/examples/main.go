package main

import (
	"fmt"
)

type Point struct{ X, Y int }

func main() {
	p := Point{1, 2}
	fmt.Println(p)

	fmt.Println(Scale(p, 5))
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}
