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
	visible := 0
	for i, line := range lines {
		array[i] = make([]int, len(line))
		for treeIndex, treeHeight := range line {
			height, _ := strconv.Atoi(string(treeHeight))
			array[i][treeIndex] = height
		}
	}

	for rowIndex, row := range array {
		for colIndex, tree := range row {
			if isVisible(rowIndex, colIndex, tree) {
				visible++
			}
		}
	}
	fmt.Println(array)
	fmt.Println(visible)
}

func isVisibleRow(rowIndex, colIndex int, height int) bool {
	visibleLeft := true
	for i := 0; i < colIndex; i++ {
		if array[rowIndex][i] >= height {
			visibleLeft = false
		}
	}
	visibleRight := true
	for i := maxCol; i > colIndex; i-- {
		if array[rowIndex][i] >= height {
			visibleRight = false
		}
	}
	return visibleLeft || visibleRight
}

func isVisibleCol(rowIndex, colIndex, height int) bool {
	visibleUp := true
	for i := 0; i < rowIndex; i++ {
		if array[i][colIndex] >= height {
			visibleUp = false
		}
	}
	visibleDown := true
	for i := maxRow; i > rowIndex; i-- {
		if array[i][colIndex] >= height {
			visibleDown = false
		}
	}
	return visibleUp || visibleDown
}

func isVisible(rowIndex, colIndex, height int) bool {
	if colIndex == 0 || colIndex == maxCol || rowIndex == 0 || rowIndex == maxRow {
		return true
	}
	return isVisibleCol(rowIndex, colIndex, height) || isVisibleRow(rowIndex, colIndex, height)
}
