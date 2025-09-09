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
