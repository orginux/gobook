package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Can't read file: %s\n%v", os.Args[1], err)
		os.Exit(1)
	}

	input := bufio.NewScanner(f)
	// Set the split function for the scanning operation.
	input.Split(bufio.ScanWords)

	words := map[string]int{}
	for input.Scan() {
		words[input.Text()]++
	}

	fmt.Printf("\nword\tcount\n")
	for word, count := range words {
		fmt.Printf("%s\t%d\n", word, count)
	}
}
