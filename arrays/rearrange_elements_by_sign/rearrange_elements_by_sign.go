package main

// You are given a 0-indexed integer array nums of even length consisting of an equal number of positive and negative integers.

// You should return the array of nums such that the the array follows the given conditions:

// Every consecutive pair of integers have opposite signs.
// For all integers with the same sign, the order in which they were present in nums is preserved.
// The rearranged array begins with a positive integer.
// Return the modified array after rearranging the elements to satisfy the aforementioned conditions.
// Constraints:

// 2 <= nums.length <= 2 * 105
// nums.length is even
// 1 <= |nums[i]| <= 105
// nums consists of equal number of positive and negative integers.

// Time: O(n) [2n], Space: O(n) [n/2+n/2]
// Intuition: Take two arrays (+ve, -ve) , iterate the original and push elements
// into their respective arrays, then loop over and modify the original
func RearrangeElementsBySignBrute(nums []int) []int {
	positives := []int{}
	negatives := []int{}
	for _, num := range nums {
		// it is said that 1 <= |nums[i]| <= 105 so we are assured num != 0
		if num > 0 {
			positives = append(positives, num)
		} else {
			negatives = append(negatives, num)
		}
	}
	for i := range len(nums) {
		if i%2 == 0 {
			nums[i] = positives[i/2]
		} else {
			nums[i] = negatives[i/2]
		}
	}
	return nums
}

// ❌Intuition: For each position figure out what should be there.
// if that is not there look in the array for that and swap them.
// NOTE: ...,1,2,3,-1,-2,-3,... suppose at the position of 1 it should be negative
// so according to our algo -1 will be replaces by it so it becomes
// ...,-1,2,3,1,-2,-3 but in doing so the order of 1,2,3, is ruined therefore it's a problem
// we can do right shift of each element to prevent the order but again that will take O(n^2)
func RearrangeElementsBySignOptimalTry(nums []int) []int {
	i := 0
	for i < len(nums) {
		if i%2 == 0 && nums[i] < 0 {
			j := i
			for j < len(nums) && nums[j] < 0 {
				j++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		if i%2 == 1 && nums[i] > 0 {
			j := i
			for j < len(nums) && nums[j] > 0 {
				j++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		i++
	}
	return nums
}

// Intuition:
func RearrangeElementsBySignOptimal(nums []int) []int {
	output := make([]int, len(nums))
	pos := 0
	neg := 1
	for i, num := range nums {
		if num > 0 {
			output[pos] = i
			pos += 2
		} else {
			output[neg] = i
			neg += 2
		}
	}
	return output
}
