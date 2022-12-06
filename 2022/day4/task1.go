package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input")
	content := strings.Split(string(file), "\n")
	total := 0
	for _, line := range content {
		elves := strings.Split(line, ",")
		elf1 := strings.Split(elves[0], "-")
		elf1low, _ := strconv.Atoi(elf1[0])
		elf1high, _ := strconv.Atoi(elf1[1])

		elf2 := strings.Split(elves[1], "-")
		elf2low, _ := strconv.Atoi(elf2[0])
		elf2high, _ := strconv.Atoi(elf2[1])

		if elf1low <= elf2low && elf1high >= elf2low || elf2low <= elf1low && elf2high >= elf1low {
			total++
			fmt.Println(elves)
		}
	}
	fmt.Println(total)
}

// 1L 2L 1H 2H
// 1L 2L 2H 1H

// 1L 1H 2L 2H
