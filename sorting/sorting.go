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
// In place sorting algorithm
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

func MergeSort(arr []int, low int, high int) {
	if low >= high {
		return
	}
	mid := int((high-low)/2) + low
	MergeSort(arr, low, mid)
	MergeSort(arr, mid+1, high)
	merge(arr, low, mid, high)
}
func merge(arr []int, low int, mid int, high int) {
	temp := []int{}
	i, j := low, mid+1
	for i <= mid && j <= high {
		if arr[i] < arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}
	for i <= mid {
		temp = append(temp, arr[i])
		i++
	}
	for j <= high {
		temp = append(temp, arr[j])
		j++
	}

	for k := low; k <= high; k++ {
		arr[k] = temp[k-low]
	}
}

func QuickSort(arr []int, low int, high int) {
	if low >= high {
		return
	}

	partIdx := partition(arr, low, high)
	QuickSort(arr, low, partIdx-1)
	QuickSort(arr, partIdx+1, high)
}
func partition(arr []int, low int, high int) int {
	pivot := arr[low]
	i := low
	j := high

	for i < j {
		// find larger elements on the right
		for i < high && arr[i] <= pivot {
			i++
		}
		// find smaller elements on the left
		for j > low && arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// swap the last element with the pivot
	arr[low], arr[j] = arr[j], arr[low]
	return j
}
