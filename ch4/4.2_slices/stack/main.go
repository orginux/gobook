package main

import "fmt"

func main() {
	var stack []int
	stack = append(stack, 1)
	stack = append(stack, 2)
	fmt.Println(stack)
	fmt.Println(stack[len(stack)-1])

	stack = stack[:len(stack)-1]
	fmt.Println(stack)
}
