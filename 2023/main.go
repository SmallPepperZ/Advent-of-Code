package main

import (
	"bufio"
	"os"
	"rutmanz/aoc/2023/days"
)

const (
	FILE_INPUT = "input.txt"
	FILE_TEST  = "input_simple.txt"
)

func loadFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return file
}
func main() {
	file := loadFile(FILE_TEST)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	day := days.Day2{}
	//day.Part1(scanner)
	day.Part2(scanner)
}
