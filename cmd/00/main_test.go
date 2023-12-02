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
			`1
      2
      3
      4
      5`,
			1,
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
			`1
      2
      3
      4
      5`,
			2,
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
