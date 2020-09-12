package main

import (
	"fmt"
	"time"
)

type Emloyee struct {
	ID       int
	Name     string
	Address  string
	DoB      time.Time
	Position string
	Salary   int
	MangerID int
}

func main() {
	var dilvert Emloyee
	fmt.Println(dilvert)

	dilvert.Salary += 5000
	position := &dilvert.Position
	*position = "Senior " + *position

	fmt.Println(dilvert)
}
