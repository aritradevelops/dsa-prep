package main

// Given an array nums containing n distinct numbers in the range [0, n],
// return the only number in the range that is missing from the array.

// Time: O(n^2)
// Intuition: For every nums 0-n check if that exists in the array
func FindMissingElementBrute(nums []int) int {
	last := len(nums)
	for i := 0; i <= last; i++ {
		exists := false
		for _, num := range nums {
			if num == i {
				exists = true
			}
		}
		if !exists {
			return i
		}
	}
	// this will not occur
	return -1
}

// Time: O(n), Space: O(n)
// Intuition: Take a visited array. Loop through the given array and mark visited
// for each number. loop through the visited array and return the not visited one
func FindMissingElementsBetter(nums []int) int {
	last := len(nums)
	visited := make([]int, last+1)
	// fmt.Println("given", nums)
	// fmt.Println("visited", visited)
	for _, num := range nums {
		visited[num] = 1
	}

	// fmt.Println("visited after", visited)
	for idx, num := range visited {
		if num == 0 {
			return idx
		}
	}
	return -1
}

// Time: O(n)
// Intuition: for first natural numbers sum is n * (n+1) / 2. for the starting zero
// this formula stays the same. therefore if we can sum all the number in the array
// and subtract it from the expected sum (i.e. n * (n+1)/2) then we can get the missing
// element
func FindMissingElementOptimal(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	last := len(nums)
	expectedSum := int(last * (last + 1) / 2)

	return expectedSum - sum
}

// Time : O(n)
// Intuition: a XOR a = 0 and 0 XOR a = a
// so for all the numbers of the given array if they are XOR'ed with
// all the numbers in range 0 - n then the matching numbers will
// become zero, and only the missing number will stay
func FindMissingElementMostOptimal(nums []int) int {
	XORed := 0
	for i, num := range nums {
		// as the ideal array will have size + 1 therefore
		// instead of 0 - len(nums), 1 - len(nums) + 1
		XORed ^= i + 1
		XORed ^= num
	}
	return XORed
}

// Why XOR version is better than sum ?
// imagine nums is 0 - 10^5 so 10^5 * (10^5 + 1) -> 10 ^10 which is really big
// but XOR(a,b) never overflows the max(a,b) therefore its slightly better
