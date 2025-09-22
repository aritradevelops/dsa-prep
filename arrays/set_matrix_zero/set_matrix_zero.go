package main

// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

// Intuition: if we directly start modifing the rows and the colums in runtime we will not be able to
// differentiate whether the zero is added by us or was in there originally, so we should store the original
// state somewhere before modifying the matrix
func SetMatrixZeroBrute(matrix [][]int) {
	copied := [][]int{}
	for _, row := range matrix {
		copiedRow := []int{}
		copiedRow = append(copiedRow, row...)
		copied = append(copied, copiedRow)
	}
	markRowZero := func(matrix [][]int, row int) {
		for i := range matrix[row] {
			matrix[row][i] = 0

		}
	}
	markColZero := func(matrix [][]int, col int) {
		for _, row := range matrix {
			row[col] = 0

		}
	}
	for i, row := range copied {
		for j, col := range row {
			if col == 0 {
				markColZero(matrix, j)
				markRowZero(matrix, i)
			}
		}
	}
}

// Intuition: we know that for a row if any element is zero then the whole row should be zero
// also for a column in any element is zero then the whole column should be zero
// so what if we keep track of whether a row contains any zero and whether a column contains any zero
// then we can definitely later mark them as zeros
func SetMatrixZeroBetter(matrix [][]int) {
	zeroedRows := make([]int, len(matrix))
	// 1 <= m, n <= 200, allows us to do access zeroth index directly
	zeroedCols := make([]int, len(matrix[0]))

	for i, row := range matrix {
		for j, col := range row {
			if col == 0 {
				zeroedCols[j] = 1
				zeroedRows[i] = 1
			}
		}
	}
	markRowZero := func(matrix [][]int, row int) {
		for i := range matrix[row] {
			matrix[row][i] = 0

		}
	}
	markColZero := func(matrix [][]int, col int) {
		for _, row := range matrix {
			row[col] = 0

		}
	}
	for i, row := range zeroedRows {
		if row == 1 {
			markRowZero(matrix, i)
		}
	}
	for j, col := range zeroedCols {
		if col == 1 {
			markColZero(matrix, j)
		}
	}
}

// A straightforward solution using O(mn) space is probably a bad idea.
// A simple improvement uses O(m + n) space, but still not the best solution.
// Could you devise a constant space solution?

// Intution: so we are using two arrays for storing the row and the column. instead of that can we use the first
// row and column to store these data? we know that if any of the row elements is zero then the full row will be zero
// so we don't need to warry about the first rows values getting lost because it will only be overwritten if it is
// destined to be zero, same goes for the colums, but if we use inded [0,0] to store the value of the first row
// then there is nothing to store the value of the first column therefore we need just another variable to store that
func SetMatrixZeroOptimal(matrix [][]int) {
	col0 := 1
	for i, row := range matrix {
		for j, col := range row {
			if col == 0 {
				matrix[i][0] = 0
				if j == 0 {
					col0 = 0
				} else {
					matrix[0][j] = 0
				}
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			// if matrix[i][0] == 0 || (j != 0 && matrix[0][j] == 0) || (j == 0 && col0 == 0) {
			// we are iterating from 1 index in both cases so we don't need explicit checking for col0
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}
	}
	if col0 == 0 {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}

}
