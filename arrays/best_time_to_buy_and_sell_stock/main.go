package main

import "github.com/aritradeveops/dsa-prep/tester"

func main() {
	t := &tester.Tester{}

	// Test cases for MaxProfitBrute
	t.Add(&tester.Test[[]int, int]{
		Name: "MaxProfitBrute",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{7, 1, 5, 3, 6, 4}, Expected: 5},
			{Input: []int{7, 6, 4, 3, 1}, Expected: 0},
			{Input: []int{1, 2}, Expected: 1},
			{Input: []int{2, 1}, Expected: 0},
			{Input: []int{1}, Expected: 0},
		},
		TestFunc: MaxProfitBrute,
	})

	// Test cases for MaxProfitOptimal
	t.Add(&tester.Test[[]int, int]{
		Name: "MaxProfitOptimal",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{7, 1, 5, 3, 6, 4}, Expected: 5},
			{Input: []int{7, 6, 4, 3, 1}, Expected: 0},
			{Input: []int{1, 2}, Expected: 1},
			{Input: []int{2, 1}, Expected: 0},
			{Input: []int{1}, Expected: 0},
		},
		TestFunc: MaxProfitOptimal,
	})

	t.Run()
}
