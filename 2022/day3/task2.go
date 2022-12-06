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
	for i := 0; i < len(content); i += 3 {
		fmt.Println(i)
		elf1 := content[i]
		elf2 := content[i+1]
		elf3 := content[i+2]
		for _, letter := range elf1 {
			if strings.Contains(elf2, string(letter)) && strings.Contains(elf3, string(letter)) {
				sum += getIndex(letter)
				break
			}
		}
	}
	fmt.Println(sum)
}
