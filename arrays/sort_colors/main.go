package main

import "github.com/aritradeveops/dsa-prep/tester"

func main() {
	t := &tester.Tester{}

	// Test cases for SortColorsBrute
	t.Add(&tester.Test[[]int, []int]{
		Name: "SortColorsBrute",
		Cases: []tester.TestCase[[]int, []int]{
			{Input: []int{}, Expected: []int{}},
			{Input: []int{2, 0, 2, 1, 1, 0}, Expected: []int{0, 0, 1, 1, 2, 2}},
			{Input: []int{2, 0, 1}, Expected: []int{0, 1, 2}},
			{Input: []int{0, 0, 0}, Expected: []int{0, 0, 0}},
			{Input: []int{2, 2, 2}, Expected: []int{2, 2, 2}},
			{Input: []int{1, 1, 1}, Expected: []int{1, 1, 1}},
			{Input: []int{0, 2, 1, 2, 0, 1, 2, 0}, Expected: []int{0, 0, 0, 1, 1, 2, 2, 2}},
		},
		TestFunc: SortColorsBruteWrapper,
	})

	// Test cases for SortColorsBetter
	t.Add(&tester.Test[[]int, []int]{
		Name: "SortColorsBetter",
		Cases: []tester.TestCase[[]int, []int]{
			{Input: []int{}, Expected: []int{}},
			{Input: []int{2, 0, 2, 1, 1, 0}, Expected: []int{0, 0, 1, 1, 2, 2}},
			{Input: []int{2, 0, 1}, Expected: []int{0, 1, 2}},
			{Input: []int{0, 0, 0}, Expected: []int{0, 0, 0}},
			{Input: []int{2, 2, 2}, Expected: []int{2, 2, 2}},
			{Input: []int{1, 1, 1}, Expected: []int{1, 1, 1}},
			{Input: []int{0, 2, 1, 2, 0, 1, 2, 0}, Expected: []int{0, 0, 0, 1, 1, 2, 2, 2}},
		},
		TestFunc: SortColorsBetterWrapper,
	})

	// Test cases for SortColorsOptimal
	t.Add(&tester.Test[[]int, []int]{
		Name: "SortColorsOptimal",
		Cases: []tester.TestCase[[]int, []int]{
			{Input: []int{}, Expected: []int{}},
			{Input: []int{2, 0, 2, 1, 1, 0}, Expected: []int{0, 0, 1, 1, 2, 2}},
			{Input: []int{2, 0, 1}, Expected: []int{0, 1, 2}},
			{Input: []int{0, 0, 0}, Expected: []int{0, 0, 0}},
			{Input: []int{2, 2, 2}, Expected: []int{2, 2, 2}},
			{Input: []int{1, 1, 1}, Expected: []int{1, 1, 1}},
			{Input: []int{0, 2, 1, 2, 0, 1, 2, 0}, Expected: []int{0, 0, 0, 1, 1, 2, 2, 2}},
		},
		TestFunc: SortColorsOptimalWrapper,
	})

	t.Run()
}

func SortColorsBruteWrapper(colors []int) []int {
	copied := make([]int, len(colors))
	copy(copied, colors)
	SortColorsBrute(copied)
	return copied
}

func SortColorsBetterWrapper(colors []int) []int {
	copied := make([]int, len(colors))
	copy(copied, colors)
	SortColorsBetter(copied)
	return copied
}

func SortColorsOptimalWrapper(colors []int) []int {
	copied := make([]int, len(colors))
	copy(copied, colors)
	SortColorsOptimal(copied)
	return copied
}
