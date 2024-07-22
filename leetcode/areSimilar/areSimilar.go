package main

import (
	"fmt"
)

func areSimilar(mat [][]int, k int) bool {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			for l := j + 1; l < len(mat[i]); l++ {
				if matrixSum(mat[i][j]) == matrixSum(mat[i][l]) {
					return true
				}
			}
		}
	}
	return false
}

func matrixSum(matrix [][]int) int {
	sum := 0
	for _, row := range matrix {
		for _, val := range row {
			sum += val
		}
	}
	return sum
}

func main() {
	slice := [][][]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}
	fmt.Println(areSimilar(slice, 4)) // Expected output: true
}
