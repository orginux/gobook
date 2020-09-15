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

	salary := &dilvert.Salary
	*salary = 1000

	Bonus(&dilvert, 20)
	fmt.Println(dilvert)
}

func Bonus(e *Emloyee, percent int) {
	e.Salary += e.Salary * percent / 100
}
