package main

import "github.com/aritradeveops/dsa-prep/tester"

type Anom struct {
	arr []int
	k   int
}

func main() {
	t := &tester.Tester{}

	// Test cases for LongestSubArrayWithSumKBrute
	t.Add(&tester.Test[Anom, int]{
		Name: "LongestSubArrayWithSumKBrute",
		Cases: []tester.TestCase[Anom, int]{
			{Input: Anom{arr: []int{}, k: 0}, Expected: 0},
			{Input: Anom{arr: []int{1, -1, 5, -2, 3}, k: 3}, Expected: 4},
			{Input: Anom{arr: []int{-2, -1, 2, 1}, k: 1}, Expected: 2},
			{Input: Anom{arr: []int{1, 2, 3}, k: 3}, Expected: 2},
			{Input: Anom{arr: []int{0, 0, 0}, k: 0}, Expected: 3},
			{Input: Anom{arr: []int{3, 4, 7, 2, -3, 1, 4, 2}, k: 7}, Expected: 4},
			{Input: Anom{arr: []int{1, 2, 1, 0, 1}, k: 3}, Expected: 3},
		},
		TestFunc: LongestSubArrayWithSumKBruteWrapper,
	})

	// Test cases for LongestSubArrayWithSumKBruteOptimized
	t.Add(&tester.Test[Anom, int]{
		Name: "LongestSubArrayWithSumKBruteOptimized",
		Cases: []tester.TestCase[Anom, int]{
			{Input: Anom{arr: []int{}, k: 0}, Expected: 0},
			{Input: Anom{arr: []int{1, -1, 5, -2, 3}, k: 3}, Expected: 4},
			{Input: Anom{arr: []int{-2, -1, 2, 1}, k: 1}, Expected: 2},
			{Input: Anom{arr: []int{1, 2, 3}, k: 3}, Expected: 2},
			{Input: Anom{arr: []int{0, 0, 0}, k: 0}, Expected: 3},
			{Input: Anom{arr: []int{3, 4, 7, 2, -3, 1, 4, 2}, k: 7}, Expected: 4},
			{Input: Anom{arr: []int{1, 2, 1, 0, 1}, k: 3}, Expected: 3},
		},
		TestFunc: LongestSubArrayWithSumKBruteOptimizedWrapper,
	})

	// Test cases for LongestSubArrayWithSumKBetter (supports negatives and zeros)
	t.Add(&tester.Test[Anom, int]{
		Name: "LongestSubArrayWithSumKBetter",
		Cases: []tester.TestCase[Anom, int]{
			{Input: Anom{arr: []int{}, k: 0}, Expected: 0},
			{Input: Anom{arr: []int{1, -1, 5, -2, 3}, k: 3}, Expected: 4},
			{Input: Anom{arr: []int{-2, -1, 2, 1}, k: 1}, Expected: 2},
			{Input: Anom{arr: []int{1, 2, 3}, k: 3}, Expected: 2},
			{Input: Anom{arr: []int{0, 0, 0}, k: 0}, Expected: 3},
			{Input: Anom{arr: []int{3, 4, 7, 2, -3, 1, 4, 2}, k: 7}, Expected: 4},
			{Input: Anom{arr: []int{1, 2, 1, 0, 1}, k: 3}, Expected: 3},
		},
		TestFunc: LongestSubArrayWithSumKBetterWrapper,
	})

	// Test cases for LongestSubArrayWithSumKOptimal (positives only, no zeros)
	t.Add(&tester.Test[Anom, int]{
		Name: "LongestSubArrayWithSumKOptimal",
		Cases: []tester.TestCase[Anom, int]{
			{Input: Anom{arr: []int{}, k: 0}, Expected: 0},
			{Input: Anom{arr: []int{1, 2, 1, 1, 1}, k: 3}, Expected: 3},
			{Input: Anom{arr: []int{1, 2, 3, 7, 5}, k: 12}, Expected: 3},
			{Input: Anom{arr: []int{1, 1, 1, 1}, k: 5}, Expected: 0},
			{Input: Anom{arr: []int{5, 1, 2, 3}, k: 5}, Expected: 2},
		},
		TestFunc: LongestSubArrayWithSumKOptimalWrapper,
	})

	t.Run()
}

func LongestSubArrayWithSumKBruteWrapper(a Anom) int {
	return LongestSubArrayWithSumKBrute(a.arr, a.k)
}

func LongestSubArrayWithSumKBruteOptimizedWrapper(a Anom) int {
	return LongestSubArrayWithSumKBruteOptimized(a.arr, a.k)
}

func LongestSubArrayWithSumKBetterWrapper(a Anom) int {
	return LongestSubArrayWithSumKBetter(a.arr, a.k)
}

func LongestSubArrayWithSumKOptimalWrapper(a Anom) int {
	return LongestSubArrayWithSumKOptimal(a.arr, a.k)
}
