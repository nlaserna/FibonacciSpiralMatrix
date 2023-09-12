package main

import (
	"fmt"
)

func convertToSpiralMatrix(sequence []int, cols, rows int) (spiralMatrix [][]int) {
	spiralMatrix = make([][]int, rows)

	for i := 0; i < len(spiralMatrix); i++ {
		spiralMatrix[i] = make([]int, cols)
	}

	if len(sequence) == 0 {
		return
	}

	top := 0
	bottom := rows - 1
	left := 0
	right := cols - 1

	index := 0

	for true {
		if left > right {
			break
		}
		for i := left; i <= right; i++ {
			spiralMatrix[top][i] = sequence[index]
			index++
		}
		top++

		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			spiralMatrix[i][right] = sequence[index]
			index++
		}
		right--

		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			spiralMatrix[bottom][i] = sequence[index]
			index++
		}
		bottom--

		if top > bottom {
			break
		}

		for i := bottom; i >= top; i-- {
			spiralMatrix[i][left] = sequence[index]
			index++
		}
		left++
	}

	return spiralMatrix
}

func SpiralMatrix(rows, cols int) (spiralMatrix [][]int) {
	n := cols * rows
	sequence := FibonacciSequence(n)

	fmt.Println("Matrix in spiral format:")
	spiralMatrix = convertToSpiralMatrix(sequence, cols, rows)
	printSpiralMatrix(spiralMatrix)
	fmt.Println()
	return spiralMatrix
}

func printSpiralMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Print("{ ")
		for j := 0; j < len(matrix[0]); j++ {
			if j == (len(matrix[0]) - 1) {
				fmt.Print(matrix[i][j])
			} else {
				fmt.Print(matrix[i][j], ", ")
			}
		}
		fmt.Print(" }")
		fmt.Println(" ")
	}
}
