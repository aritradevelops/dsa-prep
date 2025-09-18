package main

import "github.com/aritradeveops/dsa-prep/tester"

type Anom struct {
	arr []int
	k   int
}

func main() {
	t := &tester.Tester{}

	t.Add(&tester.Test[[]int, []int]{
		Name:     "Rotate Array Left By One",
		TestFunc: LeftRotateWrapper,
		Cases: []tester.TestCase[[]int, []int]{
			{Input: []int{}, Expected: []int{}},
			{Input: []int{1}, Expected: []int{1}},
			{Input: []int{1, 2}, Expected: []int{2, 1}},
			{Input: []int{0, 4, 2}, Expected: []int{4, 2, 0}},
		},
	})

	t.Add(&tester.Test[Anom, []int]{
		Name:     "Rotate Array Left By K",
		TestFunc: LeftRotateKWrapper,
		Cases: []tester.TestCase[Anom, []int]{
			{
				Input: Anom{
					arr: []int{},
					k:   5,
				},
				Expected: []int{},
			},
			{
				Input: Anom{
					arr: []int{1, 2, 3, 4, 5},
					k:   2,
				},
				Expected: []int{3, 4, 5, 1, 2},
			},
			{
				Input: Anom{
					arr: []int{1, 2, 3, 4, 5},
					k:   3,
				},
				Expected: []int{4, 5, 1, 2, 3},
			},
			{
				Input: Anom{
					arr: []int{1, 2, 3},
					k:   4,
				},
				Expected: []int{2, 3, 1},
			},
		},
	})
	t.Run()

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

	// Test cases for SingleNumberBrute
	t.Add(&tester.Test[[]int, int]{
		Name: "SingleNumberBrute",
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

	// Test cases for TwoSumBrute
	t.Add(&tester.Test[Anom, []int]{
		Name: "TwoSumBrute",
		Cases: []tester.TestCase[Anom, []int]{
			{Input: Anom{arr: []int{}, k: 0}, Expected: []int{0, 0}},
			{Input: Anom{arr: []int{2, 7, 11, 15}, k: 9}, Expected: []int{0, 1}},
			{Input: Anom{arr: []int{3, 2, 4}, k: 6}, Expected: []int{1, 2}},
			{Input: Anom{arr: []int{3, 3}, k: 6}, Expected: []int{0, 1}},
			{Input: Anom{arr: []int{-1, -2, -3, -4, -5}, k: -8}, Expected: []int{2, 4}},
			{Input: Anom{arr: []int{1, 2, 3}, k: 7}, Expected: []int{0, 0}},
		},
		TestFunc: TwoSumBruteWrapper,
	})

	// Test cases for TwoSumBetter
	t.Add(&tester.Test[Anom, []int]{
		Name: "TwoSumBetter",
		Cases: []tester.TestCase[Anom, []int]{
			{Input: Anom{arr: []int{}, k: 0}, Expected: []int{0, 0}},
			{Input: Anom{arr: []int{2, 7, 11, 15}, k: 9}, Expected: []int{1, 0}},
			{Input: Anom{arr: []int{3, 2, 4}, k: 6}, Expected: []int{2, 1}},
			{Input: Anom{arr: []int{3, 3}, k: 6}, Expected: []int{1, 0}},
			{Input: Anom{arr: []int{-1, -2, -3, -4, -5}, k: -8}, Expected: []int{4, 2}},
			{Input: Anom{arr: []int{1, 2, 3}, k: 7}, Expected: []int{0, 0}},
		},
		TestFunc: TwoSumBetterWrapper,
	})

	// Test cases for TwoSumOptimal (returns boolean)
	t.Add(&tester.Test[Anom, bool]{
		Name: "TwoSumOptimal",
		Cases: []tester.TestCase[Anom, bool]{
			{Input: Anom{arr: []int{}, k: 0}, Expected: false},
			{Input: Anom{arr: []int{2, 7, 11, 15}, k: 9}, Expected: true},
			{Input: Anom{arr: []int{3, 2, 4}, k: 6}, Expected: true},
			{Input: Anom{arr: []int{3, 3}, k: 6}, Expected: true},
			{Input: Anom{arr: []int{-1, -2, -3, -4, -5}, k: -8}, Expected: true},
			{Input: Anom{arr: []int{1, 2, 3}, k: 7}, Expected: false},
		},
		TestFunc: TwoSumOptimalWrapper,
	})

	// Run all tests
	t.Run()
}
func LeftRotateWrapper(arr []int) []int {
	copied := make([]int, len(arr))
	copy(copied, arr)
	RotateArrayLeftByOne(copied)
	return copied
}
func LeftRotateKWrapper(a Anom) []int {
	copied := make([]int, len(a.arr))
	copy(copied, a.arr)
	RotateArrayLeftByK(copied, a.k)
	return copied
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

func TwoSumBruteWrapper(a Anom) []int {
	i, j := TwoSumBrute(a.arr, a.k)
	return []int{i, j}
}

func TwoSumBetterWrapper(a Anom) []int {
	i, j := TwoSumBetter(a.arr, a.k)
	return []int{i, j}
}

func TwoSumOptimalWrapper(a Anom) bool {
	return TwoSumOptimal(a.arr, a.k)
}
