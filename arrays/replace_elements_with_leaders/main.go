package main

import "github.com/aritradeveops/dsa-prep/tester"

type arrWrap struct{ arr []int }

func main() {
	t := &tester.Tester{}

	// Tests for ReplaceWithLeadersBrute (in-place)
	t.Add(&tester.Test[arrWrap, []int]{
		Name: "ReplaceWithLeadersBrute",
		Cases: []tester.TestCase[arrWrap, []int]{
			{Input: arrWrap{arr: []int{}}, Expected: []int{}},
			{Input: arrWrap{arr: []int{17, 18, 5, 4, 6, 1}}, Expected: []int{18, 6, 6, 6, 1, -1}},
			{Input: arrWrap{arr: []int{400}}, Expected: []int{-1}},
			{Input: arrWrap{arr: []int{1, 2, 3}}, Expected: []int{3, 3, -1}},
		},
		TestFunc: func(w arrWrap) []int {
			copyArr := append([]int(nil), w.arr...)
			ReplaceWithLeadersBrute(copyArr)
			return copyArr
		},
	})

	// Tests for ReplaceWithLeadersOptimal (in-place)
	t.Add(&tester.Test[arrWrap, []int]{
		Name: "ReplaceWithLeadersOptimal",
		Cases: []tester.TestCase[arrWrap, []int]{
			{Input: arrWrap{arr: []int{}}, Expected: []int{}},
			{Input: arrWrap{arr: []int{17, 18, 5, 4, 6, 1}}, Expected: []int{18, 6, 6, 6, 1, -1}},
			{Input: arrWrap{arr: []int{400}}, Expected: []int{-1}},
			{Input: arrWrap{arr: []int{1, 2, 3}}, Expected: []int{3, 3, -1}},
		},
		TestFunc: func(w arrWrap) []int {
			copyArr := append([]int(nil), w.arr...)
			ReplaceWithLeadersOptimal(copyArr)
			return copyArr
		},
	})

	t.Run()
}
