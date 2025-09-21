package main

import "github.com/aritradeveops/dsa-prep/tester"

func main() {
	t := &tester.Tester{}

	// Test cases for SingleNumberBetter
	t.Add(&tester.Test[[]int, int]{
		Name: "SingleNumberBetter",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{2, 2, 1}, Expected: 1},
			{Input: []int{4, 1, 2, 1, 2}, Expected: 4},
			{Input: []int{1}, Expected: 1},
			{Input: []int{0, 1, 0}, Expected: 1},
			{Input: []int{-1, -1, 99}, Expected: 99},
			{Input: []int{10, 14, 14, 10, 3, 3, 7}, Expected: 7},
			{Input: []int{100000, -100000, 5, 5, 7, 7, 100000}, Expected: -100000},
		},
		TestFunc: SingleNumberBetter,
	})

	// Test cases for SingleNumberOptimal
	t.Add(&tester.Test[[]int, int]{
		Name: "SingleNumberOptimal",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{2, 2, 1}, Expected: 1},
			{Input: []int{4, 1, 2, 1, 2}, Expected: 4},
			{Input: []int{1}, Expected: 1},
			{Input: []int{0, 1, 0}, Expected: 1},
			{Input: []int{-1, -1, 99}, Expected: 99},
			{Input: []int{10, 14, 14, 10, 3, 3, 7}, Expected: 7},
			{Input: []int{100000, -100000, 5, 5, 7, 7, 100000}, Expected: -100000},
		},
		TestFunc: SingleNumberOptimal,
	})

	t.Run()
}
