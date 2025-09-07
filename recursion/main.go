package main

import "fmt"

func main() {
	PrintUptoN(0, 5)
	PrintUptoNReverse(0, 5)
	fmt.Println(SumUptoN(5))
	fmt.Println(Fact(5))
	arr := []int{1, 2, 3, 4, 5}
	ReverseArr(0, arr)
	for _, elem := range arr {
		fmt.Println(elem)
	}
	fmt.Println(IsPalindrome(0, "aritra"))
	fmt.Println(IsPalindrome(0, "madam"))
	fmt.Println(Fibonacci(4))
}

// Time Complexity: O(n) as it is n no of function calls
// Space Complexity: O(n) i.e n functions are going to wait in stack space until the nth one returns
func PrintUptoN(curr int, n int) {
	if curr == n {
		return
	}
	fmt.Println(curr + 1)
	PrintUptoN(curr+1, n)
}
func PrintUptoNReverse(curr int, n int) {
	if curr == n {
		return
	}
	PrintUptoNReverse(curr+1, n)
	fmt.Println(curr + 1)
}

func SumUptoN(n int) int {
	if n == 0 {
		return 0
	}
	return SumUptoN(n-1) + n
}
func Fact(n int) int {
	if n == 1 {
		return 1
	}
	return Fact(n-1) * n
}
func ReverseArr(idx int, arr []int) {
	opposite := len(arr) - idx - 1
	// we can also check upto idx >= n/2
	if idx >= opposite {
		return
	}
	temp := arr[idx]
	arr[idx] = arr[opposite]
	arr[opposite] = temp
	ReverseArr(idx+1, arr)
}
func IsPalindrome(idx int, str string) bool {
	opposite := len(str) - 1 - idx
	// if idx >= len(str)/2
	if idx >= opposite {
		return true
	}
	if str[opposite] != str[idx] {
		return false
	}
	return IsPalindrome(idx+1, str)
}

// For every function calls there are two more function calls
// therefore Time complexity:  2 to the power n
// to be noted that if both the call was for n-1 then it would have been accurate
// but as there is a call for n-2 the time complexity will be slightly lesser
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
