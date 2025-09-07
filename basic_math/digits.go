package main

import (
	"fmt"
	"math"
	"sort"
)

// Time Complexity : O(log10(n))
func CountDigitOne(num int) int {
	count := 0
	for num > 0 {
		// lastDigit := num % 10
		num = (num / 10)
		count++
	}
	return count
}
func CountDigitTwo(num int) int {
	return int(math.Log10(float64(num))) + 1
}

func ReverseNumber(num int) int {
	reversed := 0
	for num > 0 {
		lastDigit := num % 10
		reversed = reversed*10 + lastDigit
		num = int(num / 10)
	}
	return reversed
}

func CheckPalindrome(num int) bool {
	return num == ReverseNumber(num)
}

func IsArmstrongNumber(num int) bool {
	count := CountDigitTwo(num)
	calculated := 0
	original := num
	for num > 0 {
		lastDigit := num % 10
		num = int(num / 10)
		calculated += int(math.Pow(float64(lastDigit), float64(count)))
	}
	return original == calculated
}

func GetFactors(num int) []int {
	output := []int{}
	// instead of sqrt(num) we can do i * i which is comuptationaly cheaper
	// we should not loop from 0 because it will cause DivideByZero Err
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			output = append(output, i)
			other := int(num / i)
			if i != other {
				output = append(output, other)
			}
		}
	}
	// Just for the test cases
	sort.Ints(output)
	return output
}

// if a number only have 2 factors that is 1 and the number itself
func IsPrime(num int) bool {
	count := 0
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			count++
			other := int(num / count)
			if other != i {
				count++
			}
		}
	}
	fmt.Printf("No : %d, Count: %d", num, count)
	return count == 2
}

// GCD = Greates Common Divisor | HCF = Highest Common Factor
func FindGCD(a int, b int) int {
	smaller := int(math.Min(float64(a), float64(b)))
	gcd := 0
	for i := smaller; i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			gcd = i
			break
		}
	}
	return gcd
}

// Euclidean Algorighm : GCD(a,b) = GCD(a-b, b) where a > b
// instead of doing a - b everytime we can do a%b to substract b,the no of
// times it would have, at once, when one of the no becomes 0 then the
// other number is the GCD
// or else we can say
// GCD(a,b) = {a if b = 0 | GCD(b, a mod b) otherwise} where we have to maintain a > b always manually
// Time Complexity: O(logϕ(min(a,b)))
// Why ϕ ? because a and b is getting divided by arbitary no's so we don't know exactly
func FindGCDUsingEuclideanlAlgorithm(a int, b int) int {
	for a != 0 && b != 0 {
		// if a is greater then a % b
		if a > b {
			a = a % b
		} else {
			// else for lower or equal value b % a, equal is the thing to note here
			b = b % a
		}
	}
	if a == 0 {
		return b
	}
	return a
}

// LCM = Least Common Multiple
func FindLCM(a int, b int) int {
	return (a * b) / FindGCDUsingEuclideanlAlgorithm(a, b)
}
