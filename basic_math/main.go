package main

import "github.com/aritradeveops/dsa-prep/tester"

func main() {
	t := &tester.Tester{}
	t.Add(&tester.Test[int, int]{
		Name:     "Count Digits",
		TestFunc: CountDigitOne,
		Cases: []tester.TestCase[int, int]{
			{
				Input:    5120,
				Expected: 4,
			},
			{
				Input:    200,
				Expected: 3,
			},
			{
				Input:    123456789,
				Expected: 9,
			},
		},
	})

	// Two-int input wrapper
	type Pair struct{ A, B int }

	// FindGCD tests
	t.Add(&tester.Test[Pair, int]{
		Name:     "GCD (Brute Force)",
		TestFunc: func(p Pair) int { return FindGCD(p.A, p.B) },
		Cases: []tester.TestCase[Pair, int]{
			{Input: Pair{12, 18}, Expected: 6},
			{Input: Pair{100, 10}, Expected: 10},
			{Input: Pair{7, 7}, Expected: 7},
			{Input: Pair{17, 13}, Expected: 1},
			{Input: Pair{17, 0}, Expected: 0},
		},
	})

	// FindGCDUsingEuclideanlAlgorithm tests
	t.Add(&tester.Test[Pair, int]{
		Name:     "GCD (Euclidean)",
		TestFunc: func(p Pair) int { return FindGCDUsingEuclideanlAlgorithm(p.A, p.B) },
		Cases: []tester.TestCase[Pair, int]{
			{Input: Pair{12, 18}, Expected: 6},
			{Input: Pair{100, 10}, Expected: 10},
			{Input: Pair{7, 7}, Expected: 7},
			{Input: Pair{17, 13}, Expected: 1},
		},
	})

	// FindLCM tests
	t.Add(&tester.Test[Pair, int]{
		Name:     "LCM",
		TestFunc: func(p Pair) int { return FindLCM(p.A, p.B) },
		Cases: []tester.TestCase[Pair, int]{
			{Input: Pair{4, 6}, Expected: 12},
			{Input: Pair{21, 6}, Expected: 42},
			{Input: Pair{5, 7}, Expected: 35},
			{Input: Pair{10, 10}, Expected: 10},
		},
	})

	// CountDigitTwo tests
	t.Add(&tester.Test[int, int]{
		Name:     "Count Digits (Log10)",
		TestFunc: CountDigitTwo,
		Cases: []tester.TestCase[int, int]{
			{Input: 5120, Expected: 4},
			{Input: 200, Expected: 3},
			{Input: 123456789, Expected: 9},
		},
	})

	// ReverseNumber tests
	t.Add(&tester.Test[int, int]{
		Name:     "Reverse Number",
		TestFunc: ReverseNumber,
		Cases: []tester.TestCase[int, int]{
			{Input: 123, Expected: 321},
			{Input: 120, Expected: 21},
			{Input: 5, Expected: 5},
			{Input: 1000, Expected: 1},
		},
	})

	// CheckPalindrome tests
	t.Add(&tester.Test[int, bool]{
		Name:     "Check Palindrome",
		TestFunc: CheckPalindrome,
		Cases: []tester.TestCase[int, bool]{
			{Input: 121, Expected: true},
			{Input: 123, Expected: false},
			{Input: 0, Expected: true},
			{Input: 1221, Expected: true},
		},
	})

	// IsArmstrongNumber tests
	t.Add(&tester.Test[int, bool]{
		Name:     "Armstrong Number",
		TestFunc: IsArmstrongNumber,
		Cases: []tester.TestCase[int, bool]{
			{Input: 153, Expected: true},
			{Input: 370, Expected: true},
			{Input: 371, Expected: true},
			{Input: 407, Expected: true},
			{Input: 100, Expected: false},
			{Input: 10, Expected: false},
		},
	})

	// Get Factors tests
	t.Add(&tester.Test[int, []int]{
		Name:     "Get Factors",
		TestFunc: GetFactors,
		Cases: []tester.TestCase[int, []int]{
			{Input: 56, Expected: []int{1, 2, 4, 7, 8, 14, 28, 56}},
			{Input: 2, Expected: []int{1, 2}},
			{Input: 4, Expected: []int{1, 2, 4}},
			{Input: 36, Expected: []int{1, 2, 3, 4, 6, 9, 12, 18, 36}},
		},
	})
	// Is Prime tests
	t.Add(&tester.Test[int, bool]{
		Name:     "Is Prime",
		TestFunc: IsPrime,
		Cases: []tester.TestCase[int, bool]{
			{Input: 0, Expected: false},
			{Input: 1, Expected: false},
			{Input: 2, Expected: true},
			{Input: 4, Expected: false},
			{Input: 7, Expected: true},
			{Input: 9, Expected: false},
		},
	})
	t.Run()
}
