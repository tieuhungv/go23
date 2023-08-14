package main

import (
	"fmt"
)

func countRectangles(rectangles [][]int) int {
	counts := 0
	rows := len(rectangles)
	columns := len(rectangles[0])

	for indexRow := 0; indexRow < rows; indexRow++ {
		for indexCol := 0; indexCol < columns; indexCol++ {
			nextRow := indexRow + 1
			nextCol := indexCol + 1
			prevCol := indexCol - 1
			if rectangles[indexRow][indexCol] == 1 {
				isLeftVal := nextRow == rows-1 || (prevCol == -1 || rectangles[indexRow][prevCol] == 0)
				isRightVal := nextCol == columns || rectangles[indexRow][nextCol] == 0
				isBottVal := nextRow == rows || rectangles[nextRow][indexCol] == 0
				if !isBottVal {
					for index := indexCol; index < columns; index++ {
						if nextRow != rows && rectangles[indexRow][index] == 1 {
							rectangles[indexRow][index] = 0
						} else {
							break
						}
					}
				}
				if isLeftVal && isRightVal && isBottVal {
					fmt.Println(indexRow, indexCol)
					counts++
				}
			}
		}
	}

	return counts
}

func main() {
	arr := [][]int{
		{0, 0, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0},
		{0, 0, 1, 1, 1, 0},
		{0, 0, 1, 0, 0, 1},
		{0, 0, 1, 1, 0, 0},
		{0, 1, 0, 0, 1, 0},
	}

	count := countRectangles(arr)
	fmt.Println(count)
}
