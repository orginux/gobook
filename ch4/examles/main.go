package main

import "fmt"

func main() {
	s := []string{"one", "", "three", "", "five"}
	s = nonemty(s)
	fmt.Println(s)
}

func nonemty(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
