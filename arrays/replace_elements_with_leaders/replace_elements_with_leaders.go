package main

// Given an array arr, replace every element in that array with the greatest element among
// the elements to its right, and replace the last element with -1.

// After doing so, return the array.

// Time: O(n^2)
// Intuition: For each element find the greater to its right and replace it with that
func ReplaceWithLeadersBrute(nums []int) {
	for i := range nums {
		leader := -1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > leader {
				leader = nums[j]
			}
		}
		nums[i] = leader
	}
}

// Time: O(n)
// Intuition: iterate from right to left and keep track of the greatest at the right
// replace the current element with the greatest(leader) and check if current number is greater
// then it becomes the next leader
// contraints : 1 <= arr.length <= 104
// 1 <= arr[i] <= 105, if arr had negative numbers(-2 but leader still would have been -1)
// then starting with - 1 won't work
func ReplaceWithLeadersOptimal(nums []int) {
	leader := -1
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		nums[i] = leader
		if num > leader {
			leader = num
		}
	}
}
