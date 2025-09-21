package main

import "github.com/aritradeveops/dsa-prep/tester"

func main() {
	t := &tester.Tester{}

	// Test cases for MajorityElementBrute
	t.Add(&tester.Test[[]int, int]{
		Name: "MajorityElementBrute",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{3}, Expected: 3},
			{Input: []int{3, 3, 4}, Expected: 3},
			{Input: []int{2, 2, 1, 1, 1, 2, 2}, Expected: 2},
			{Input: []int{1, 1, 1, 2, 2}, Expected: 1},
			{Input: []int{6, 5, 5, 5, 6, 5, 5}, Expected: 5},
			{Input: []int{9, 9, 9, 9, 8, 8}, Expected: 9},
		},
		TestFunc: MajorityElementBrute,
	})

	// Test cases for MajorityElementBetter
	t.Add(&tester.Test[[]int, int]{
		Name: "MajorityElementBetter",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{3}, Expected: 3},
			{Input: []int{3, 3, 4}, Expected: 3},
			{Input: []int{2, 2, 1, 1, 1, 2, 2}, Expected: 2},
			{Input: []int{1, 1, 1, 2, 2}, Expected: 1},
			{Input: []int{6, 5, 5, 5, 6, 5, 5}, Expected: 5},
			{Input: []int{9, 9, 9, 9, 8, 8}, Expected: 9},
		},
		TestFunc: MajorityElementBetter,
	})

	// Test cases for MajorityElementOptimal
	t.Add(&tester.Test[[]int, int]{
		Name: "MajorityElementOptimal",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{3}, Expected: 3},
			{Input: []int{3, 3, 4}, Expected: 3},
			{Input: []int{2, 2, 1, 1, 1, 2, 2}, Expected: 2},
			{Input: []int{1, 1, 1, 2, 2}, Expected: 1},
			{Input: []int{6, 5, 5, 5, 6, 5, 5}, Expected: 5},
			{Input: []int{9, 9, 9, 9, 8, 8}, Expected: 9},
		},
		TestFunc: MajorityElementOptimal,
	})

	t.Run()
}
