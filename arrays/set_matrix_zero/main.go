package main

import "github.com/aritradeveops/dsa-prep/tester"

type matrixWrap struct{ m [][]int }

func main() {
	t := &tester.Tester{}

	// Brute tests
	t.Add(&tester.Test[matrixWrap, [][]int]{
		Name: "SetMatrixZeroBrute",
		Cases: []tester.TestCase[matrixWrap, [][]int]{
			{Input: matrixWrap{m: [][]int{}}, Expected: [][]int{}},
			{Input: matrixWrap{m: [][]int{{1}}}, Expected: [][]int{{1}}},
			{Input: matrixWrap{m: [][]int{{0}}}, Expected: [][]int{{0}}},
			{Input: matrixWrap{m: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}}, Expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
			{Input: matrixWrap{m: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}}, Expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		},
		TestFunc: func(w matrixWrap) [][]int {
			copyM := deepCopy(w.m)
			SetMatrixZeroBrute(copyM)
			return copyM
		},
	})

	// Better tests
	t.Add(&tester.Test[matrixWrap, [][]int]{
		Name: "SetMatrixZeroBetter",
		Cases: []tester.TestCase[matrixWrap, [][]int]{
			{Input: matrixWrap{m: [][]int{}}, Expected: [][]int{}},
			{Input: matrixWrap{m: [][]int{{1}}}, Expected: [][]int{{1}}},
			{Input: matrixWrap{m: [][]int{{0}}}, Expected: [][]int{{0}}},
			{Input: matrixWrap{m: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}}, Expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
			{Input: matrixWrap{m: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}}, Expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		},
		TestFunc: func(w matrixWrap) [][]int {
			copyM := deepCopy(w.m)
			SetMatrixZeroBetter(copyM)
			return copyM
		},
	})

	// Optimal tests
	t.Add(&tester.Test[matrixWrap, [][]int]{
		Name: "SetMatrixZeroOptimal",
		Cases: []tester.TestCase[matrixWrap, [][]int]{
			{Input: matrixWrap{m: [][]int{}}, Expected: [][]int{}},
			{Input: matrixWrap{m: [][]int{{1}}}, Expected: [][]int{{1}}},
			{Input: matrixWrap{m: [][]int{{0}}}, Expected: [][]int{{0}}},
			{Input: matrixWrap{m: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}}, Expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
			{Input: matrixWrap{m: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}}, Expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		},
		TestFunc: func(w matrixWrap) [][]int {
			copyM := deepCopy(w.m)
			SetMatrixZeroOptimal(copyM)
			return copyM
		},
	})

	t.Run()
}

func deepCopy(m [][]int) [][]int {
	res := make([][]int, len(m))
	for i := range m {
		res[i] = append([]int(nil), m[i]...)
	}
	return res
}
