package main

import "sort"

// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.
// You can return the answer in any order.

// Time: O(n^2) [n * (n+1) / 2]
// Intuition: For every two element in the array check if they add upto the given number
func TwoSumBrute(nums []int, k int) (int, int) {
	for i, num1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			if (num1 + nums[j]) == k {
				return i, j
			}
		}
	}
	return 0, 0
}

// Time: O(n), Space: O(n)
// Intuition: Store the other number in an map
// for any val check if k - val exists in the map
func TwoSumBetter(nums []int, k int) (int, int) {
	visitedMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := visitedMap[k-num]; ok {
			return i, j
		}
		visitedMap[num] = i
	}
	return 0, 0
}

// Intuition: Sort the array and use two pointer to adjust the value to k
// indexes can not be returned only boolean
func TwoSumOptimal(nums []int, k int) bool {
	sort.Ints(nums)
	i, j := 0, len(nums)-1

	for i < j {
		sum := nums[i] + nums[j]
		if sum == k {
			return true
		}
		if sum < k {
			i++
		} else {
			j--
		}
	}
	return false
}
