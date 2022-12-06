package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	content := string(file)
	lines := strings.Split(content, "\n")
	elves := make([]int, len(lines))
	elfindex := 0
	for _, line := range lines {
		value, _ := strconv.Atoi(line)
		elves[elfindex] += value

		if line == "" {
			elfindex++
		}
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})
	fmt.Println(elves[0]+elves[1]+elves[2])
}
