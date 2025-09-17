package main

// Given a binary array nums, return the maximum number of consecutive 1's in the array.

// Time: O(n^2)
// Intuition: for every element in the array if its an one
// the no of consecutive ones will be counted and if that's
// greater than the maximum then maximum will be updated
func MaximumConsecutiveOnesBrute(nums []int) int {
	maximum := 0

	for i, num := range nums {
		if num == 0 {
			continue
		}
		count := 0
		for j := i; j < len(nums); j++ {
			if nums[j] == 1 {
				count++
			} else {
				break
			}
		}
		if count > maximum {
			maximum = count
		}
		// now for the 1's that are in the right of this
		// counting the maximum is useless as the longest sequence
		// is already been calculated here, so we can start from the
		// next zero and that will be our optimal solution
	}
	return maximum
}

// Time: O(n)
// Intuition: For every new 1 the consecutive ones will be counted
// if that's more than the maximum then maximum will be updated
// and this will be repeated until the last

func MaximumConsecutiveOnesOptimalBadCode(nums []int) int {
	maximum := 0
	start := 0

	for start < len(nums) {
		// finding the next one
		for start < len(nums) && nums[start] == 0 {
			start++
		}
		count := 0
		// count the consecutive ones
		for start < len(nums) {
			if nums[start] == 0 {
				break
			}
			count++
			start++
		}
		if count > maximum {
			maximum = count
		}

	}
	// all i'm looping over is start
	return maximum
}

// Intuition: If i encounter a zero i will reset the count
// else for every one i will increase the count and update the maximum
func MaximumConsecutiveOnesOptimalGoodCode(nums []int) int {
	maximum := 0
	count := 0
	for _, num := range nums {
		if num == 0 {
			count = 0
			continue
		}
		count++
		if count > maximum {
			maximum = count
		}
	}
	return maximum
}
