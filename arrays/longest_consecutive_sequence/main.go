package main

import "github.com/aritradeveops/dsa-prep/tester"

func main() {
	t := &tester.Tester{}

	// Tests for LongestConsecutiveSequenceBrute
	t.Add(&tester.Test[[]int, int]{
		Name: "LongestConsecutiveSequenceBrute",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{}, Expected: 0},
			{Input: []int{100, 4, 200, 1, 3, 2}, Expected: 4},
			{Input: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, Expected: 9},
			{Input: []int{1, 2, 0, 1}, Expected: 3},
		},
		TestFunc: LongestConsecutiveSequenceBrute,
	})

	// Tests for LongestConsecutiveSequenceOptimal
	t.Add(&tester.Test[[]int, int]{
		Name: "LongestConsecutiveSequenceOptimal",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{}, Expected: 0},
			{Input: []int{100, 4, 200, 1, 3, 2}, Expected: 4},
			{Input: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, Expected: 9},
			{Input: []int{1, 2, 0, 1}, Expected: 3},
		},
		TestFunc: LongestConsecutiveSequenceOptimal,
	})

	t.Run()
}
