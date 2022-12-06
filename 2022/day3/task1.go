package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getIndex(letter rune) int {
	return strings.Index(alphabet, string(letter)) + 1
}

func main() {
	file, _ := ioutil.ReadFile("input")
	content := strings.Split(string(file), "\n")
	sum := 0
	for _, line := range content {
		if line == "" {
			break
		}
		left := line[:len(line)/2]
		right := line[len(line)/2:]
		for _, letter := range left {
			if strings.Contains(right, string(letter)) {
				sum += getIndex(letter)
				break
			}
		}
	}
	fmt.Println(sum)
}
