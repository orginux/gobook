package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var result385 = flag.Bool("sha384", false, "Use sha384")
	var result512 = flag.Bool("sha512", false, "Use sha512")
	flag.Parse()

	args := flag.Args()
	if len(args) > 1 {
		log.Fatal("Use one Arg")
		os.Exit(1)
	}

	if *result385 {
		sum384 := sha512.Sum384([]byte(args[0]))
		fmt.Printf("%x\n", sum384)
	} else if *result512 {
		sum512 := sha512.Sum512([]byte(args[0]))
		fmt.Printf("%x\n", sum512)
	} else {
		sum256 := sha256.Sum256([]byte(args[0]))
		fmt.Printf("%x\n", sum256)
	}
}
