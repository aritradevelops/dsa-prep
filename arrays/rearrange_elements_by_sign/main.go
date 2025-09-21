package main

import "github.com/aritradeveops/dsa-prep/tester"

func main() {
	t := &tester.Tester{}

	// Test cases for RearrangeElementsBySignBrute
	t.Add(&tester.Test[[]int, []int]{
		Name: "RearrangeElementsBySignBrute",
		Cases: []tester.TestCase[[]int, []int]{
			{Input: []int{3, 1, -2, -5, 2, -4}, Expected: []int{3, -2, 1, -5, 2, -4}},
			{Input: []int{-1, 1}, Expected: []int{1, -1}},
			{Input: []int{1, 2, 3, -1, -2, -3}, Expected: []int{1, -1, 2, -2, 3, -3}},
		},
		TestFunc: RearrangeElementsBySignBruteWrapper,
	})

	t.Run()
}

func RearrangeElementsBySignBruteWrapper(nums []int) []int {
	copied := make([]int, len(nums))
	copy(copied, nums)
	return RearrangeElementsBySignBrute(copied)
}
