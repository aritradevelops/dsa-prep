package main

// Given an array arr[] of size n containing integers, the task is to find the length of the longest subarray having sum equal to the given value k.

// Note: If there is no subarray with sum equal to k, return 0.

// Time: O(n^3)
// Intuition: generate all the subarrays
// find the ones with sum K
// return the longest length
func LongestSubArrayWithSumKBrute(nums []int, k int) int {
	longest := 0
	for i := range nums {
		for j := i; j < len(nums); j++ {
			sum := 0
			// Time: O(n^3)
			// now instead of calculating some to the end every time
			// we could use sum calculated upto previous element and
			// add the current element to get the sum of the subarray
			// from i to j
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			if sum == k {
				subArrLen := j - i + 1
				if subArrLen > longest {
					longest = subArrLen
				}
			}
		}
	}
	return longest
}

// Time: O(n^2)
// Intuition: we are optimizing the inner loop from the previous observations
func LongestSubArrayWithSumKBruteOptimized(nums []int, k int) int {
	longest := 0
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			// sum of the subarray i to j
			sum += nums[j]
			if sum == k {
				length := j - i + 1
				if length > longest {
					longest = length
				}
			}
		}
	}
	return longest
}

// This is the most optimal for nums containing zeros and negative values
// Time: O(n), Space: O(n)
// Intuition: starring at ith element if sum = x, and there exits a
// previous some which is x - k at jth index, then the some of the elements
// from j to i is x - (x - k) = k
func LongestSubArrayWithSumKBetter(nums []int, k int) int {
	prefixMap := make(map[int]int)
	prefixSum := 0
	longest := 0
	for i, num := range nums {
		prefixSum += num
		if prefixSum == k {
			// then it is the longest
			longest = i + 1
		} else if val, ok := prefixMap[(prefixSum - k)]; ok {
			length := i - val
			if length > longest {
				longest = length
			}
		}
		// If this prefixSum exists previously then do not overwrite it
		// as we want the length to be the longest
		if _, ok := prefixMap[prefixSum]; !ok {
			prefixMap[prefixSum] = i
		}
	}
	return longest
}

// Only valid for positives numbers and not zeros
// Time : O(n) [2n]
// Intuition: Sliding Window
// the sum can be increased while it is less than k
// and it can be decreased while it is greater than k to find k
func LongestSubArrayWithSumKOptimal(nums []int, k int) int {
	windowStart, windowEnd, sum, longest := 0, 0, 0, 0
	if k == 0 {
		return longest
	}
	for windowEnd < len(nums) {
		sum += nums[windowEnd]

		for sum > k {
			sum -= nums[windowStart]
			windowStart++
		}
		if sum == k {
			length := windowEnd - windowStart + 1
			if length > longest {
				longest = length
			}
		}
		windowEnd++
	}
	return longest
}
