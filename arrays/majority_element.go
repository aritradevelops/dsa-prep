package main

// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.

// Time: O(n^2)
// Intuition: For each element i will count the number of occurrences in that array
// and if that exists more that n/2 times then return that
func MajorityElementBrute(nums []int) int {
	for _, num := range nums {
		count := 0
		for _, num2 := range nums {
			if num == num2 {
				count++
			}
		}
		if count > (len(nums) / 2) {
			return num
		}
	}
	// this will not occur
	return -1
}

// Time: O(n) Space: O(n) [n/2]
// Intuition: keep track of the occurrences and return whenever it becomes more that n / 2
// 2, 2, 1, 1, 1, 2, 2
func MajorityElementBetter(nums []int) int {
	occurrences := make(map[int]int)
	for _, num := range nums {
		total := 0
		if occur, ok := occurrences[num]; ok {
			total = occur + 1
		} else {
			total = 1
		}
		occurrences[num] = total
		if total > (len(nums) / 2) {
			return num
		}
	}
	// this will not occur
	return -1
}

// Time: O(n)
// Intuition:
// When the elements are the same as the candidate element, votes are incremented whereas
// when some other element is found (not equal to the candidate element), we decreased the count.
// This actually means that we are decreasing the priority of winning ability of the selected candidate,
// since we know that if the candidate is in majority it occurs more than N/2 times and the remaining elements
// are less than N/2. We keep decreasing the votes since we found some different element(s) than the candidate element.
// When votes become 0, this actually means that there are the equal  number of votes for different elements,
// which should not be the case for the element to be the majority element. So the candidate element cannot be
// the majority and hence we choose the present element as the candidate and continue the same till all the
// elements get finished. The final candidate would be our majority element. We check using the 2nd traversal
// to see whether its count is greater than N/2. If it is true, we consider it as the majority element.
func MajorityElementOptimal(nums []int) int {
	candidate := nums[0]
	count := 1
	for j := 1; j < len(nums); j++ {
		if count == 0 {
			candidate = nums[j]
		}
		if nums[j] == candidate {
			count++
		} else {
			count--
		}
	}
	// if we are not assured that there exists a majority element
	// then we would need to check and confirm with another iteration
	return candidate
}
