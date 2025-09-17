package main

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// You must implement a solution with a linear runtime complexity and use only constant extra space.

// Brute: for each element linear search for its pair, Time: O(n^2)

// Time: O(n) [n +( n/2+1)] Space: O(n) [n/2+1]
// Map is used for long numbers like 10^12 and negatives
func SingleNumberBetter(nums []int) int {
	occurrences := make(map[int]int)

	for _, num := range nums {
		if val, ok := occurrences[num]; ok {
			occurrences[num] = val + 1
		} else {
			occurrences[num] = 1
		}
	}

	for k, v := range occurrences {
		if v == 1 {
			return k
		}
	}
	return -1
}

// Time : O(n)
// Intuition: a XOR a = 0 and 0 XOR a = a
// so all the pairs will be canceled out and only the single will left
func SingleNumberOptimal(nums []int) int {
	XORed := 0
	for _, num := range nums {
		XORed ^= num
	}
	return XORed
}

// Some other approaches are:
// sort and traverse to find single element
// add them in a set, sum all the elements in the set, multiply it by two,
// subtract sum of all the elements from this and divide by two
