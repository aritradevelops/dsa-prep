package main

import "github.com/aritradeveops/dsa-prep/tester"

func main() {
	t := &tester.Tester{}

	// Test cases for FindMissingElementBrute
	t.Add(&tester.Test[[]int, int]{
		Name: "FindMissingElementBrute",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{3, 0, 1}, Expected: 2},
			{Input: []int{0, 1}, Expected: 2},
			{Input: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, Expected: 8},
			{Input: []int{1, 2, 3, 4, 5}, Expected: 0},
			{Input: []int{0, 1, 2, 3, 4}, Expected: 5},
			{Input: []int{0}, Expected: 1},
			{Input: []int{1}, Expected: 0},
			{Input: []int{0, 1, 2, 3, 4, 5, 6, 7, 9}, Expected: 8},
		},
		TestFunc: FindMissingElementBrute,
	})

	// Test cases for FindMissingElementsBetter
	t.Add(&tester.Test[[]int, int]{
		Name: "FindMissingElementsBetter",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{3, 0, 1}, Expected: 2},
			{Input: []int{0, 1}, Expected: 2},
			{Input: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, Expected: 8},
			{Input: []int{1, 2, 3, 4, 5}, Expected: 0},
			{Input: []int{0, 1, 2, 3, 4}, Expected: 5},
			{Input: []int{0}, Expected: 1},
			{Input: []int{1}, Expected: 0},
			{Input: []int{0, 1, 2, 3, 4, 5, 6, 7, 9}, Expected: 8},
		},
		TestFunc: FindMissingElementsBetter,
	})

	// Test cases for FindMissingElementOptimal
	t.Add(&tester.Test[[]int, int]{
		Name: "FindMissingElementOptimal",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{3, 0, 1}, Expected: 2},
			{Input: []int{0, 1}, Expected: 2},
			{Input: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, Expected: 8},
			{Input: []int{1, 2, 3, 4, 5}, Expected: 0},
			{Input: []int{0, 1, 2, 3, 4}, Expected: 5},
			{Input: []int{0}, Expected: 1},
			{Input: []int{1}, Expected: 0},
			{Input: []int{0, 1, 2, 3, 4, 5, 6, 7, 9}, Expected: 8},
		},
		TestFunc: FindMissingElementOptimal,
	})

	// Test cases for FindMissingElementMostOptimal
	t.Add(&tester.Test[[]int, int]{
		Name: "FindMissingElementMostOptimal",
		Cases: []tester.TestCase[[]int, int]{
			{Input: []int{3, 0, 1}, Expected: 2},
			{Input: []int{0, 1}, Expected: 2},
			{Input: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, Expected: 8},
			{Input: []int{1, 2, 3, 4, 5}, Expected: 0},
			{Input: []int{0, 1, 2, 3, 4}, Expected: 5},
			{Input: []int{0}, Expected: 1},
			{Input: []int{1}, Expected: 0},
			{Input: []int{0, 1, 2, 3, 4, 5, 6, 7, 9}, Expected: 8},
		},
		TestFunc: FindMissingElementMostOptimal,
	})

	t.Run()
}
