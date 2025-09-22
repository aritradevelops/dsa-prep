package main

// Intuition: to traverse i know i have to move to the right first
// then to the bottom, then to the left, then to the top
func SpiralTraversal(matrix [][]int) []int {
	result := []int{}
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for left <= right && top <= bottom {
		// left to right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		// top to bottom
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		// only if there exists more element to print
		if top <= bottom {
			// right to left
			for i := right; i > left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}
		// only if there exists more element to print
		if left <= right {
			for i := bottom; i > top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
		// bottom to top
	}
	return result
}
