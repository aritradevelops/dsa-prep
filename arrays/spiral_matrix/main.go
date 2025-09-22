package main

import "fmt"

func main() {
	// matrix := [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
	matrix2 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	fmt.Println(SpiralTraversal(matrix2))
}
