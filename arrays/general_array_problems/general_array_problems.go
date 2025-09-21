package main

import "fmt"

func RotateArrayLeftByOne(arr []int) {
	if len(arr) > 0 {
		first := arr[0]
		for i := 1; i <= len(arr)-1; i++ {
			arr[i-1] = arr[i]
		}
		arr[len(arr)-1] = first
	}
}

// 1, 2, 3, 4, 5 | 2
func RotateArrayLeftByK(arr []int, k int) {
	if len(arr) > 0 {
		k = k % len(arr)
		if k != 0 {
			firstElements := append([]int(nil), arr[0:k]...)
			// fmt.Println(firstElements, len(firstElements))
			for i := k; i <= len(arr)-1; i++ {
				arr[i-k] = arr[i]
			}
			// fmt.Println("here", arr, firstElements)
			for j := 0; j < len(firstElements); j++ {
				arr[len(arr)-k+j] = firstElements[j]
			}
			// fmt.Println("here after", arr)
		}
	}
}

func PutZerosAtTheEnd(nums []int) {
	i := 0
	for i < len(nums) && nums[i] != 0 {
		i++
	}
	for j := i + 1; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			// for nums[i] != 0 {
			// 	i++
			// }
			// instead of the above we can simply do i++ because
			// we are asure that either there was some zero's in between
			// else it is just the next element I switched with which should
			// currently contain the zero that I swapped it with
			i++
		}
	}
}

func RemoveDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		// if they are the same
		// just don't care because
		// when a new guy will come it will be replaced by them
		if nums[i] == nums[j] {
			continue
		}
		// ok so the new guy has came to take place
		nums[i+1], nums[j] = nums[j], nums[i+1]
		// now i + 1 contains a new element
		// so we point i to it so that it can be compared with the incremented
		// j at next iteration
		i++
	}
	return i + 1
}

func RemoveDuplicatesII(nums []int) {
	i := 0
	count := 1
	for j := 1; j < len(nums); j++ {
		if nums[i] == nums[j] {
			count++
			if count <= 2 {
				if j-i != 1 {
					nums[i+1], nums[j] = nums[j], nums[i+1]
				}
				i++
			}
		} else {
			count = 1
			nums[i+1], nums[j] = nums[j], nums[i+1]
			i++
		}
	}
}

func LongestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix
	}
	for i, ch := range strs[0] {
		flag := true
		fmt.Println(i, string(ch))
		for j := 1; j < len(strs); j++ {
			fmt.Println("===>", rune(strs[j][i]), rune(ch))
			if len(strs[j]) == i {
				flag = false
				break
			}
			if rune(strs[j][i]) != ch {
				flag = false
				break
			}
		}
		fmt.Println(flag)
		if !flag {
			break
		}
		prefix = prefix + string(ch)
	}
	return prefix
}
