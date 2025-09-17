package main

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
