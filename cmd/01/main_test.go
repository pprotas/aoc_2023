package main

import "testing"

func TestPartOne(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			"Part One",
			`1abc2
      pqr3stu8vwx
      a1b2c3d4e5f
      treb7uchet`,
			142,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := partOne(tc.input)
			if actual != tc.expected {
				t.Errorf("partOne(%v) = %v, expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			"Part Two",
			`two1nine
      eightwothree
      abcone2threexyz
      xtwone3four
      4nineeightseven2
      zoneight234
      7pqrstsixteen`,
			281,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := partTwo(tc.input)
			if actual != tc.expected {
				t.Errorf("partTwo(%v) = %v, expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}
