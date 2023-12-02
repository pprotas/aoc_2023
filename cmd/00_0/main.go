package main

import (
	"fmt"
	"log"

	"aoc_2023/internal/input"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	input, err := input.GetInput(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)
}
