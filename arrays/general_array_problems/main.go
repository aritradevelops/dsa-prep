package main

import "github.com/aritradeveops/dsa-prep/tester"

type Anom struct {
	arr []int
	k   int
}

func main() {
	t := &tester.Tester{}

	t.Add(&tester.Test[[]int, []int]{
		Name:     "Rotate Array Left By One",
		TestFunc: LeftRotateWrapper,
		Cases: []tester.TestCase[[]int, []int]{
			{Input: []int{}, Expected: []int{}},
			{Input: []int{1}, Expected: []int{1}},
			{Input: []int{1, 2}, Expected: []int{2, 1}},
			{Input: []int{0, 4, 2}, Expected: []int{4, 2, 0}},
		},
	})

	t.Add(&tester.Test[Anom, []int]{
		Name:     "Rotate Array Left By K",
		TestFunc: LeftRotateKWrapper,
		Cases: []tester.TestCase[Anom, []int]{
			{
				Input: Anom{
					arr: []int{},
					k:   5,
				},
				Expected: []int{},
			},
			{
				Input: Anom{
					arr: []int{1, 2, 3, 4, 5},
					k:   2,
				},
				Expected: []int{3, 4, 5, 1, 2},
			},
			{
				Input: Anom{
					arr: []int{1, 2, 3, 4, 5},
					k:   3,
				},
				Expected: []int{4, 5, 1, 2, 3},
			},
			{
				Input: Anom{
					arr: []int{1, 2, 3},
					k:   4,
				},
				Expected: []int{2, 3, 1},
			},
		},
	})
	t.Add(&tester.Test[[]int, []int]{
		Name:     "Put all zeros at the end",
		TestFunc: PutZerosAtTheEndWrapper,
		Cases: []tester.TestCase[[]int, []int]{
			// mixed case
			{Input: []int{1, 2, 0, 0, 1, 0, 2, 3, 0, 0, 0, 5}, Expected: []int{1, 2, 1, 2, 3, 5, 0, 0, 0, 0, 0, 0}},
			// all zeros
			{Input: []int{0, 0, 0}, Expected: []int{0, 0, 0}},
			// no zeros
			{Input: []int{1, 2, 3}, Expected: []int{1, 2, 3}},
			// zeros in front
			{Input: []int{0, 0, 1, 2}, Expected: []int{1, 2, 0, 0}},
			// single element
			{Input: []int{0}, Expected: []int{0}},
			{Input: []int{5}, Expected: []int{5}},
		},
	})
	t.Add(&tester.Test[[]int, []int]{
		Name:     "Remove Duplicates",
		TestFunc: RemoveDuplicatesWrapper,
		Cases: []tester.TestCase[[]int, []int]{
			// empty
			// {Input: []int{}, Expected: []int{}},
			// single element
			{Input: []int{1}, Expected: []int{1}},
			// no duplicates
			{Input: []int{1, 2, 3}, Expected: []int{1, 2, 3}},
			// consecutive duplicates
			{Input: []int{1, 1, 2, 2, 3, 3}, Expected: []int{1, 2, 3}},
			// all same
			{Input: []int{5, 5, 5, 5}, Expected: []int{5}},
		},
	})

	t.Run()
}

func LeftRotateWrapper(arr []int) []int {
	copied := make([]int, len(arr))
	copy(copied, arr)
	RotateArrayLeftByOne(copied)
	return copied
}

func LeftRotateKWrapper(a Anom) []int {
	copied := make([]int, len(a.arr))
	copy(copied, a.arr)
	RotateArrayLeftByK(copied, a.k)
	return copied
}

func PutZerosAtTheEndWrapper(arr []int) []int {
	copied := make([]int, len(arr))
	copy(copied, arr)
	PutZerosAtTheEnd(copied)
	return copied
}

func RemoveDuplicatesWrapper(arr []int) []int {
	copied := make([]int, len(arr))
	copy(copied, arr)
	uniqueElems := RemoveDuplicates(copied)
	return copied[0:uniqueElems]
}
