package main

import "math"

// Given an integer array nums, find the subarray with the largest sum, and return its sum.

// Time: O(n^2)
// Intuition: Generate all the subarrays and sum them up to find the largest value
func MaximumSubArraySumBrute(nums []int) int {
	maxi := math.MinInt
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > maxi {
				maxi = sum
			}
		}
	}
	return maxi
}

// Time: O(n)
// Intuition: Kadane's algorithm
// we will keep adding up until the prefix sum is less than zero
func MaximumSubArraySumOptimal(nums []int) int {
	maxi := math.MinInt
	sum := 0
	for _, num := range nums {
		sum += num
		// set the maximum even if it is in minus
		if maxi < sum {
			maxi = sum
		}
		// then reset it to zero
		if sum < 0 {
			sum = 0
		}
	}
	return maxi
}
