package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var points = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}
var play_translate = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

func main() {
	file, _ := ioutil.ReadFile("input")
	content := strings.Split(string(file), "\n")
	score := 0
	for _, line := range content {
		if line == "" {
			break
		}
		opponent_raw, you := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		opponent := play_translate[opponent_raw]
		score += points[you]
		if opponent == you {
			score += 3
		} else if (opponent == "X" && you == "Y") || (opponent == "Y" && you == "Z") || (opponent == "Z" && you == "X") {
			score += 6
		}
	}
	fmt.Println(score)
}
