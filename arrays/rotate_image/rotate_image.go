package main

import "slices"

// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
// DO NOT allocate another 2D matrix and do the rotation.

// Intuition: ith row becomes (n-i-1)th column therefore what if we can make a new matrix
// rotate over each row in reverse order and start adding it in rows
func RotateImageBrute(image [][]int) {
	newImage := make([][]int, len(image))
	for i := range newImage {
		newImage[i] = make([]int, len(image))
	}

	for i := len(image) - 1; i >= 0; i-- {
		row := image[i]
		for cIdx, col := range row {
			newImage[cIdx][(len(image) - i - 1)] = col
		}
	}
	for i, row := range newImage {
		copy(image[i], row)
	}
}

// Intuition: From observation we can see starting from the bottom left of the original array
// if we move to the top then one right, then to the bottom and then repeat and in the duplicate if we
// start filling up starting from the top left to the right, then one bottom then to the left and continue
// we can have the matrix rotated, will implement later.
// Another Intuition: iterating through the original matrix in spiral direction and shifting each element by the
// length of the matrix - 1 times gives the updated matrix

// Optimal Intuition: if we observe the rotated output we can see that from top left to bottom right corner (\) is
// rotating and becoming top right to bottom left corner. to achieve this we can transpose the matrix and then reverse
// each row and this 90 deg rotation is achieved by 180 deg 3d rotation around the corner and a 180 deg 2d rotation
// w.r.to the middle axis

func RotateImageOptimal(matrix [][]int) {
	// transpose
	for i := range matrix {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// reverse all row
	for _, row := range matrix {
		slices.Reverse(row)
	}
}
