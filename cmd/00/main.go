package main

import (
	"aoc_2023/internal/input"
	"flag"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	partTwoFlag := flag.Bool("parttwo", false, "Execute part two")
	flag.Parse()

	input, err := input.GetInput(1)
	if err != nil {
		log.Fatal(err)
	}

	if *partTwoFlag {
		log.Println("Part two: ", partTwo(input))
	} else {
		log.Println("Part one: ", partOne(input))
	}
}

func partOne(input string) int {
	return 1
}

func partTwo(input string) int {
	return 2
}
