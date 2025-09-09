package main

import "github.com/aritradeveops/dsa-prep/tester"

func main() {
	t := &tester.Tester{}

	// SelectionSort tests
	t.Add(&tester.Test[[]int, []int]{
		Name:     "SelectionSort",
		TestFunc: SelectionSortWrapper,
		Cases: []tester.TestCase[[]int, []int]{
			{Input: []int{3, 2, 5, 7, 1}, Expected: []int{1, 2, 3, 5, 7}},
			{Input: []int{5, 4, 3, 2, 1}, Expected: []int{1, 2, 3, 4, 5}},
			{Input: []int{1, 2, 3, 4, 5}, Expected: []int{1, 2, 3, 4, 5}},
			{Input: []int{2, 2, 2}, Expected: []int{2, 2, 2}},
			{Input: []int{}, Expected: []int{}},
		},
	})

	// BubbleSort tests
	t.Add(&tester.Test[[]int, []int]{
		Name:     "BubbleSort",
		TestFunc: BubbleSortWrapper,
		Cases: []tester.TestCase[[]int, []int]{
			{Input: []int{3, 2, 5, 7, 1}, Expected: []int{1, 2, 3, 5, 7}},
			{Input: []int{5, 4, 3, 2, 1}, Expected: []int{1, 2, 3, 4, 5}},
			{Input: []int{1, 2, 3, 4, 5}, Expected: []int{1, 2, 3, 4, 5}},
			{Input: []int{2, 2, 2}, Expected: []int{2, 2, 2}},
			{Input: []int{}, Expected: []int{}},
		},
	})
	// InsertionSort tests
	t.Add(&tester.Test[[]int, []int]{
		Name:     "InsertionSort",
		TestFunc: InsertionSortWrapper,
		Cases: []tester.TestCase[[]int, []int]{
			{Input: []int{3, 2, 5, 7, 1}, Expected: []int{1, 2, 3, 5, 7}},
			{Input: []int{5, 4, 3, 2, 1}, Expected: []int{1, 2, 3, 4, 5}},
			{Input: []int{1, 2, 3, 4, 5}, Expected: []int{1, 2, 3, 4, 5}},
			{Input: []int{2, 2, 2}, Expected: []int{2, 2, 2}},
			{Input: []int{}, Expected: []int{}},
		},
	})
	// MergeSort tests
	t.Add(&tester.Test[[]int, []int]{
		Name:     "MergeSort",
		TestFunc: MergeSortWrapper,
		Cases: []tester.TestCase[[]int, []int]{
			{Input: []int{3, 2, 5, 7, 1}, Expected: []int{1, 2, 3, 5, 7}},
			{Input: []int{5, 4, 3, 2, 1}, Expected: []int{1, 2, 3, 4, 5}},
			{Input: []int{1, 2, 3, 4, 5}, Expected: []int{1, 2, 3, 4, 5}},
			{Input: []int{2, 2, 2}, Expected: []int{2, 2, 2}},
			{Input: []int{}, Expected: []int{}},
		},
	})
	// QuickSort tests
	t.Add(&tester.Test[[]int, []int]{
		Name:     "QuickSort",
		TestFunc: QuickSortWrapper,
		Cases: []tester.TestCase[[]int, []int]{
			{Input: []int{3, 2, 5, 7, 1}, Expected: []int{1, 2, 3, 5, 7}},
			{Input: []int{5, 4, 3, 2, 1}, Expected: []int{1, 2, 3, 4, 5}},
			{Input: []int{1, 2, 3, 4, 5}, Expected: []int{1, 2, 3, 4, 5}},
			{Input: []int{2, 2, 2}, Expected: []int{2, 2, 2}},
			{Input: []int{}, Expected: []int{}},
		},
	})

	t.Run()
}

// Wrappers to adapt in-place sorters to tester's input/output shape
func SelectionSortWrapper(input []int) []int {
	copied := make([]int, len(input))
	copy(copied, input)
	SelectionSort(copied)
	return copied
}

func BubbleSortWrapper(input []int) []int {
	copied := make([]int, len(input))
	copy(copied, input)
	BubbleSort(copied)
	return copied
}
func InsertionSortWrapper(input []int) []int {
	copied := make([]int, len(input))
	copy(copied, input)
	InsertionSort(copied)
	return copied
}
func MergeSortWrapper(input []int) []int {
	copied := make([]int, len(input))
	copy(copied, input)
	MergeSort(copied, 0, len(copied)-1)
	return copied
}
func QuickSortWrapper(input []int) []int {
	copied := make([]int, len(input))
	copy(copied, input)
	QuickSort(copied, 0, len(copied)-1)
	return copied
}
