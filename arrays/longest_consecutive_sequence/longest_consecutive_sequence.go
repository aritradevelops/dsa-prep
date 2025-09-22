package main

import (
	"slices"
)

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// Time: O(n^2)
func LongestConsecutiveSequenceBrute(nums []int) int {
	longest := 0
	for _, num := range nums {
		count := 1
		for next := num + 1; ; next++ {
			idx := slices.Index(nums, next)
			if idx != -1 {
				count++
			} else {
				break
			}
		}
		if count > longest {
			longest = count
		}
	}
	return longest
}

func LongestConsecutiveSequenceOptimal(nums []int) int {
	visited := make(map[int]struct{})
	for _, num := range nums {
		visited[num] = struct{}{}
	}
	longest := 0
	for k := range visited {
		// if we have previous element of this number
		// then counting from this is a waste of time
		if _, ok := visited[k-1]; ok {
			continue
		}
		count := 1
		for next := k + 1; ; next++ {
			if _, ok := visited[next]; !ok {
				break
			}
			count++
		}
		if count > longest {
			longest = count
		}
	}
	return longest
}
