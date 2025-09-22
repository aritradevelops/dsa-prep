package main

func SubArraySumEqualsK(nums []int, k int) int {
	prefixMap := make(map[int]int)
	prefixMap[0] = 1
	sum := 0
	count := 0
	for _, num := range nums {
		sum += num
		remaining := sum - k
		if freq, ok := prefixMap[remaining]; ok {
			count = count + freq
		}
		if prefixSum, ok := prefixMap[sum]; ok {
			prefixMap[sum] = prefixSum + 1
		} else {
			prefixMap[sum] = 1
		}
	}
	return count
}
