package main

import "github.com/aritradeveops/dsa-prep/tester"

type Anom struct {
	arr []int
	k   int
}

func main() {
	t := &tester.Tester{}
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
	t.Run()
}
func TwoSumBetterWrapper(a Anom) []int {
	i, j := TwoSumBetter(a.arr, a.k)
	return []int{i, j}
}

func TwoSumOptimalWrapper(a Anom) bool {
	return TwoSumOptimal(a.arr, a.k)
}
func TwoSumBruteWrapper(a Anom) []int {
	i, j := TwoSumBrute(a.arr, a.k)
	return []int{i, j}
}
