package main

import "github.com/aritradeveops/dsa-prep/tester"

func main() {
	t := &tester.Tester{}

	// Test cases for MaximumConsecutiveOnesBrute
	t.Add(&tester.Test[[]int, int]{
		Name: "MaximumConsecutiveOnesBrute",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{1, 1, 0, 1, 1, 1}, Expected: 3},
			{Input: []int{1, 0, 1, 1, 0, 1}, Expected: 2},
			{Input: []int{0, 0, 0, 0}, Expected: 0},
			{Input: []int{1, 1, 1, 1, 1}, Expected: 5},
			{Input: []int{1}, Expected: 1},
			{Input: []int{0}, Expected: 0},
			{Input: []int{}, Expected: 0},
			{Input: []int{1, 0, 1, 0, 1, 0, 1}, Expected: 1},
			{Input: []int{0, 1, 1, 1, 0, 1, 1, 0}, Expected: 3},
			{Input: []int{1, 1, 1, 0, 0, 1, 1, 1, 1, 1}, Expected: 5},
		},
		TestFunc: MaximumConsecutiveOnesBrute,
	})

	// Test cases for MaximumConsecutiveOnesOptimalBadCode
	t.Add(&tester.Test[[]int, int]{
		Name: "MaximumConsecutiveOnesOptimalBadCode",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{1, 1, 0, 1, 1, 1}, Expected: 3},
			{Input: []int{1, 0, 1, 1, 0, 1}, Expected: 2},
			{Input: []int{0, 0, 0, 0}, Expected: 0},
			{Input: []int{1, 1, 1, 1, 1}, Expected: 5},
			{Input: []int{1}, Expected: 1},
			{Input: []int{0}, Expected: 0},
			{Input: []int{}, Expected: 0},
			{Input: []int{1, 0, 1, 0, 1, 0, 1}, Expected: 1},
			{Input: []int{0, 1, 1, 1, 0, 1, 1, 0}, Expected: 3},
			{Input: []int{1, 1, 1, 0, 0, 1, 1, 1, 1, 1}, Expected: 5},
		},
		TestFunc: MaximumConsecutiveOnesOptimalBadCode,
	})

	// Test cases for MaximumConsecutiveOnesOptimalGoodCode
	t.Add(&tester.Test[[]int, int]{
		Name: "MaximumConsecutiveOnesOptimalGoodCode",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{1, 1, 0, 1, 1, 1}, Expected: 3},
			{Input: []int{1, 0, 1, 1, 0, 1}, Expected: 2},
			{Input: []int{0, 0, 0, 0}, Expected: 0},
			{Input: []int{1, 1, 1, 1, 1}, Expected: 5},
			{Input: []int{1}, Expected: 1},
			{Input: []int{0}, Expected: 0},
			{Input: []int{}, Expected: 0},
			{Input: []int{1, 0, 1, 0, 1, 0, 1}, Expected: 1},
			{Input: []int{0, 1, 1, 1, 0, 1, 1, 0}, Expected: 3},
			{Input: []int{1, 1, 1, 0, 0, 1, 1, 1, 1, 1}, Expected: 5},
		},
		TestFunc: MaximumConsecutiveOnesOptimalGoodCode,
	})

	t.Run()
}
