package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"releqsed"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
