package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)       // count Unicode chars
	var utflen [utf8.UTFMax + 1]int    // count lens UTF-8
	invalid := 0                       // count invalid chars UTF-8
	categories := make(map[string]int) // count categories

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // return rune, bytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r) {
			categories["letters"]++
		}
		if unicode.IsNumber(r) {
			categories["numbers"]++
		}
		if unicode.IsSpace(r) {
			categories["spaces"]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	fmt.Printf("\ncategory\tcount\n")
	for l, n := range categories {
		fmt.Printf("%s\t\t%d\n", l, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d incorrect chars UTF-8\n", invalid)
	}
}
