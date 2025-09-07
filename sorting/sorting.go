package main

// Intution: Select the minimum and sort
// Time Complexity: O(n^2)
// In place sorting algorithm
func SelectionSort(arr []int) {
	n := len(arr)
	// we need 1 less iteration as if we have sorted n-1
	// elements then the nth element will be automatically sorted
	for i := 0; i < n-1; i++ {
		min_idx := i
		// the ith element is considered the minimum
		// therefore we can start checking from i + 1
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min_idx] {
				min_idx = j
			}
		}
		arr[i], arr[min_idx] = arr[min_idx], arr[i]
	}
}

// Intuition: Continuously swap lower elements with higher elements
// Time Complexity: O(n^2), Best Case: O(n)
// In place sorting algorighm
func BubbleSort(arr []int) {
	n := len(arr)
	// we need 1 less iteration as if we have sorted n-1
	// elements then the nth element will be automatically sorted
	for i := 0; i < n-1; i++ {
		swapped := false
		// we need the loop upto the second last element as
		// as we are comparing it with the last element (j+1)
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

// At ith iteration subarray[0,i] should be sorted
// if not sorted then sort it by swapping element until
// it's in its correct position
func InsertionSort(arr []int) {
	length := len(arr)
	// we can skip the 0th iteration as it has no element on left to compare
	for i := 1; i < length; i++ {
		j := i
		// swap the bigger elements to the right
		// until everything is sorted upto that point
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}
