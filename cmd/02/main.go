package main

import (
	"aoc_2023/internal/input"
	"bufio"
	"flag"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	partTwoFlag := flag.Bool("parttwo", false, "Execute part two")
	flag.Parse()

	input, err := input.GetInput(2)
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
	const maxRed = 12
	const maxGreen = 13
	const maxBlue = 14

	scanner := bufio.NewScanner(strings.NewReader(input))

	sum := 0
next:
	for scanner.Scan() {
		line := scanner.Text()

		idRegex := regexp.MustCompile(`^Game (\d+):`)
		id, err := strconv.Atoi(idRegex.FindStringSubmatch(line)[1])
		if err != nil {
			log.Fatal(err)
		}

		redRegex := regexp.MustCompile(`(\d+) red`)
		redMatches := redRegex.FindAllStringSubmatch(line, -1)
		if len(redMatches) > 0 {
			for _, match := range redMatches {
				amount, err := strconv.Atoi(match[1])
				if err != nil {
					log.Fatal(err)
				}
				if amount > maxRed {
					continue next
				}
			}
		}
		blueRegex := regexp.MustCompile(`(\d+) blue`)
		blueMatches := blueRegex.FindAllStringSubmatch(line, -1)
		if len(blueMatches) > 0 {
			for _, match := range blueMatches {
				amount, err := strconv.Atoi(match[1])
				if err != nil {
					log.Fatal(err)
				}
				if amount > maxBlue {
					continue next
				}
			}
		}

		greenRegex := regexp.MustCompile(`(\d+) green`)
		greenMatches := greenRegex.FindAllStringSubmatch(line, -1)
		if len(greenMatches) > 0 {
			for _, match := range greenMatches {
				amount, err := strconv.Atoi(match[1])
				if err != nil {
					log.Fatal(err)
				}
				if amount > maxGreen {
					continue next
				}
			}
		}

		sum += id
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}

	return sum
}

func partTwo(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	sumPower := 0
	for scanner.Scan() {
		smallestRed := 0
		smallestBlue := 0
		smallestGreen := 0

		line := scanner.Text()

		redRegex := regexp.MustCompile(`(\d+) red`)
		redMatches := redRegex.FindAllStringSubmatch(line, -1)
		if len(redMatches) > 0 {
			for _, match := range redMatches {
				amount, err := strconv.Atoi(match[1])
				if err != nil {
					log.Fatal(err)
				}
				if amount > smallestRed {
					smallestRed = amount
				}
			}
		}

		blueRegex := regexp.MustCompile(`(\d+) blue`)
		blueMatches := blueRegex.FindAllStringSubmatch(line, -1)
		if len(blueMatches) > 0 {
			for _, match := range blueMatches {
				amount, err := strconv.Atoi(match[1])
				if err != nil {
					log.Fatal(err)
				}
				if amount > smallestBlue {
					smallestBlue = amount
				}
			}
		}

		greenRegex := regexp.MustCompile(`(\d+) green`)
		greenMatches := greenRegex.FindAllStringSubmatch(line, -1)
		if len(greenMatches) > 0 {
			for _, match := range greenMatches {
				amount, err := strconv.Atoi(match[1])
				if err != nil {
					log.Fatal(err)
				}
				if amount > smallestGreen {
					smallestGreen = amount
				}
			}
		}

		sumPower += smallestRed * smallestBlue * smallestGreen
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}

	return sumPower
}
