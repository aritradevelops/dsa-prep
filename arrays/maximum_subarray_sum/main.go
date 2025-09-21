package main

import "github.com/aritradeveops/dsa-prep/tester"

func main() {
	t := &tester.Tester{}

	// Test cases for MaximumSubArraySumBrute
	t.Add(&tester.Test[[]int, int]{
		Name: "MaximumSubArraySumBrute",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{1}, Expected: 1},
			{Input: []int{-1}, Expected: -1},
			{Input: []int{1, 2, 3}, Expected: 6},
			{Input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, Expected: 6},
			{Input: []int{5, 4, -1, 7, 8}, Expected: 23},
			{Input: []int{-2, -3, -1}, Expected: -1},
			{Input: []int{0, 0, 0}, Expected: 0},
			{Input: []int{100, -100, 50}, Expected: 100},
			{Input: []int{4, -1, 2, 1}, Expected: 6},
		},
		TestFunc: MaximumSubArraySumBrute,
	})

	// Test cases for MaximumSubArraySumOptimal
	t.Add(&tester.Test[[]int, int]{
		Name: "MaximumSubArraySumOptimal",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{1}, Expected: 1},
			{Input: []int{-1}, Expected: -1},
			{Input: []int{1, 2, 3}, Expected: 6},
			{Input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, Expected: 6},
			{Input: []int{5, 4, -1, 7, 8}, Expected: 23},
			{Input: []int{-2, -3, -1}, Expected: -1},
			{Input: []int{0, 0, 0}, Expected: 0},
			{Input: []int{100, -100, 50}, Expected: 100},
			{Input: []int{4, -1, 2, 1}, Expected: 6},
		},
		TestFunc: MaximumSubArraySumOptimal,
	})

	t.Run()
}
