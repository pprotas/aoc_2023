package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"aoc_2023/internal/input"
)

var mapping = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

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

func getIndexAndNumber(input string) (int, string) {
	lowestIndex := len(input)
	theNumber := ""
	for number := range mapping {
		index := strings.Index(input, number)
		if index < lowestIndex && index != -1 {
			lowestIndex = index
			theNumber = number
		}
	}

	return lowestIndex, theNumber
}

func wordsToNumbers(input string) string {
	for {
		index, number := getIndexAndNumber(input)
		if index == len(input) {
			break
		}
		input = input[:index+1] + mapping[number] + input[index+len(number)-1:]
	}
	return input
}

func execute(input string, parse func(input string) string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	total := 0
	for scanner.Scan() {
		line := parse(scanner.Text())
		combined, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit(line), lastDigit(line)))
		if err != nil {
			log.Fatal(err)
		}
		total += combined
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}

	return total
}

func partOne(input string) int {
	return execute(input, func(line string) string {
		return line
	})
}

func partTwo(input string) int {
	return execute(input, func(line string) string {
		return wordsToNumbers(line)
	})
}

func firstDigit(input string) int {
	for _, char := range input {
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
	}
	return 0
}

func lastDigit(input string) int {
	return firstDigit(reverseString(input))
}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
