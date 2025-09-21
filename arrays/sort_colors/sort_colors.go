package main

import "sort"

// Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
// You must solve this problem without using the library's sort function.

// Time: O(NlogN)
// Intuition: using existing sorting methods
func SortColorsBrute(colors []int) {
	sort.Ints(colors)
}

// Time: O(N)[2N] Space: O(3)
// Intuition: count occurrences then put them manually
func SortColorsBetter(colors []int) {
	occurrences := make(map[int]int)
	for _, color := range colors {
		if occur, ok := occurrences[color]; ok {
			occurrences[color] = occur + 1
		} else {
			occurrences[color] = 1
		}
	}
	occur0 := occurrences[0]
	occur1 := occurrences[1]
	occur2 := occurrences[2]

	for i := range occur0 {
		colors[i] = 0
	}
	for j := occur0; j < occur0+occur1; j++ {
		colors[j] = 1
	}
	for k := occur0 + occur1; k < occur0+occur1+occur2; k++ {
		colors[k] = 2
	}
}

// The Dutch national flag problem[1] is a computational problem proposed by Edsger Dijkstra.[2]
// The flag of the Netherlands consists of three colors: red, white, and blue.
// Given balls of these three colors arranged randomly in a line (it does not matter how many balls there are),
// the task is to arrange them such that all balls of the same color are together and their collective color
// groups are in the correct order.
// Intuition:
// rule 1 : 0...(low-1) should contain only 0s
// rule 2 : low...(mid-1) should contain only 1s
// rule 3 : mid...high is unsorted
// rule 4 : high+1...n-1 should contain only 2s
// we will iterate through the unsorted portion and try to put them into their correction position
// initially the whole array is our unsorted portion so low = 0, mid=0, high = n -1
// if mid == 0 => upto low-1 should be 0 so we can swap them and then increase low and mid
// if mid == 1 => upto mid-1 should be 1 so we can increase mid to the next position
// if mid == 2 => from high + 1 should be 2 so we can swap it with high and decrease low to maintain that
func SortColorsOptimal(colors []int) {
	low := 0
	mid := 0
	high := len(colors) - 1
	for mid <= high {
		switch colors[mid] {
		case 0:
			colors[low], colors[mid] = colors[mid], colors[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			colors[mid], colors[high] = colors[high], colors[mid]
			high--
		}
	}
}
