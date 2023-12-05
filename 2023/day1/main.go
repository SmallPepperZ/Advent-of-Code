package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	findNext := func(arr []byte) int {
		if len(arr) >= 3 {
			strToCheck := string(arr[:3])
			switch strToCheck {
			case "one":
				return 1
			case "two":
				return 2
			case "six":
				return 6
			}
		}

		if len(arr) >= 4 {
			strToCheck := string(arr[:4])
			switch strToCheck {
			case "four":
				return 4
			case "five":
				return 5
			case "nine":
				return 9
			}
		}
		if len(arr) >= 5 {
			strToCheck := string(arr[:5])
			switch strToCheck {
			case "three":
				return 3
			case "eight":
				return 8
			case "seven":
				return 7
			}
		}
		return 0
	}
	for scanner.Scan() {
		text := scanner.Bytes()
		first := -1
		last := -1
		for i, char := range text {

			num, err := strconv.Atoi(string(char))
			written := findNext(text[i:])
			if err == nil {
				if first == -1 {
					first = num
				}
				last = num
			}
			if written != 0 {
				if first == -1 {
					first = written
				}
				last = written
			}
		}
		sum += 10*first + last
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
