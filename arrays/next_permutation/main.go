package main

import "github.com/aritradeveops/dsa-prep/tester"

func main() {
	t := &tester.Tester{}

	// Test cases for NextPermutationOptimal
	t.Add(&tester.Test[[]int, []int]{
		Name: "NextPermutationOptimal",
		Cases: []tester.TestCase[[]int, []int]{
			{Input: []int{1, 2, 3}, Expected: []int{1, 3, 2}},
			{Input: []int{3, 2, 1}, Expected: []int{1, 2, 3}},
			{Input: []int{1, 1, 5}, Expected: []int{1, 5, 1}},
			{Input: []int{1}, Expected: []int{1}},
			{Input: []int{1, 2}, Expected: []int{2, 1}},
			{Input: []int{2, 1, 5, 3, 0, 0}, Expected: []int{2, 3, 0, 0, 1, 5}},
		},
		TestFunc: NextPermutationOptimalWrapper,
	})

	t.Run()
}

func NextPermutationOptimalWrapper(nums []int) []int {
	copied := make([]int, len(nums))
	copy(copied, nums)
	NextPermutationOptimal(copied)
	return copied
}
