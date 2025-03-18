package main

import "fmt"

func main() {
	matrix := [3][4]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}

	var zeroRows, zeroCols []int

	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				zeroRows = append(zeroRows, i)
				zeroCols = append(zeroCols, j)
			}
		}
	}
	for _, row := range zeroRows {
		for j := 0; j < len(matrix[row]); j++ {
			matrix[row][j] = 0
		}
	}

	for _, col := range zeroCols {
		for i := 0; i < len(matrix); i++ {
			matrix[i][col] = 0
		}
	}
	for _, row := range matrix {
		fmt.Println(row)
	}

	var arrnums = [9]int{1, 5, 4, 4, 1, 8, 8, 7, 5}
	n := len(arrnums)
	for i := 0; i < n; i++ {
		unique := true

		for j := 0; j < n; j++ {
			if i != j && arrnums[i] == arrnums[j] {
				unique = false
				break
			}
		}
		if unique {
			println(arrnums[i])

		}

	}

}
