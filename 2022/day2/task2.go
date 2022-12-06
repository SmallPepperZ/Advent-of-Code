package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var points = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}
var play_points = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

var winners = map[string]string{
	"A": "B",
	"B": "C",
	"C": "A",
}

func getLosingPlay(toLose string) string {
	for play, winner := range winners {
		if winner == toLose {
			return play
		}
	}
	return ""
}

func main() {
	file, _ := ioutil.ReadFile("input")
	content := strings.Split(string(file), "\n")
	score := 0
	for _, line := range content {
		if line == "" {
			break
		}
		opponent, result := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		score += points[result]

		if result == "X" {
			score += play_points[getLosingPlay(opponent)]
		} else if result == "Y" {
			score += play_points[opponent]
		} else {
			score += play_points[winners[opponent]]
		}
	}
	fmt.Println(score)
}
