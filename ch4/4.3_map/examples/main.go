package main

import "fmt"

func main() {
	a := map[string]bool{
		"A": true,
		"B": false,
	}
	for key, value := range a {
		fmt.Printf("%s\t%t\n", key, value)
	}
}
