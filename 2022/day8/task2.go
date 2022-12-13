package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var array [][]int
var maxCol, maxRow int

func main() {
	file, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(file), "\n")
	array = make([][]int, len(lines))
	maxRow = len(lines) - 1
	maxCol = len(lines[0]) - 1

	for i, line := range lines {
		array[i] = make([]int, len(line))
		for treeIndex, treeHeight := range line {
			height, _ := strconv.Atoi(string(treeHeight))
			array[i][treeIndex] = height
		}
	}
	bestScore := 0
	for rowIndex, row := range array {
		for colIndex, tree := range row {
			score := getScore(rowIndex, colIndex, tree)
			if score > bestScore {
				bestScore = score
			}

		}
	}
	fmt.Println(array)
	fmt.Println(bestScore)
}

func rowScore(rowIndex, colIndex int, height int) (int, int) {
	visibleLeft := 0
	for i := colIndex - 1; i >= 0; i-- {
		fmt.Println(colIndex, i, height, array[rowIndex][i])
		if array[rowIndex][i] < height {
			visibleLeft++
		} else {
			visibleLeft++
			break
		}
	}
	visibleRight := 0
	for i := colIndex + 1; i <= maxCol; i++ {
		
		if array[rowIndex][i] < height {
			visibleRight++
		} else {
			visibleRight++
			break
		}
	}
	return visibleLeft, visibleRight
}

func colScore(rowIndex, colIndex, height int) (int, int) {
	visibleUp := 0
	for i := rowIndex - 1; i >= 0; i-- {
		if array[i][colIndex] < height {
			visibleUp++
		} else {
			visibleUp++
			break
		}
	}
	visibleDown := 0
	for i := rowIndex + 1; i <= maxRow; i++ {
		if array[i][colIndex] < height {
			visibleDown++
		} else {
			visibleDown++
			break
		}
	}
	return visibleUp, visibleDown
}

func getScore(rowIndex, colIndex, height int) int {
	if colIndex == 0 || colIndex == maxCol || rowIndex == 0 || rowIndex == maxRow {
		return 0
	}
	col1, col2 := colScore(rowIndex, colIndex, height)
	row1, row2 := rowScore(rowIndex, colIndex, height)
	fmt.Println(col1*col2*row1*row2, "[", rowIndex, colIndex, "]", col1, col2, row1, row2)

	return col1 * col2 * row1 * row2
}
