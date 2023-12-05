package days

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type Day2 struct{}

func (_ Day2) Part2(scanner *bufio.Scanner) {
	checkCubes := func(color string, text string) int {
		regex, _ := regexp.Compile("(\\d+) " + color)
		matches := regex.FindAllStringSubmatch(text, -1)
		maxVal := 0
		for _, match := range matches {
			num, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			if num > maxVal {
				maxVal = num
			}
		}
		return maxVal
	}
	out := 0
	for scanner.Scan() {
		//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		text := scanner.Text()
		idRegex, _ := regexp.Compile("Game (\\d+): ")
		idmatch := idRegex.FindStringSubmatch(text)
		id, _ := strconv.Atoi(idmatch[1])
		text = text[len(idmatch[0]):]

		blues := checkCubes("blue", text)
		reds := checkCubes("red", text)
		greens := checkCubes("green", text)
		power := blues * reds * greens
		fmt.Println(id, power)
		out += power
	}
	fmt.Println(out)
}

func (_ Day2) Part1(scanner *bufio.Scanner) {
	checkCubes := func(color string, max int, text string) bool {
		regex, _ := regexp.Compile("(\\d+) " + color)
		matches := regex.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			num, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			if num > max {
				return false
			}
		}
		return true
	}
	out := 0
	for scanner.Scan() {
		//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		text := scanner.Text()
		idRegex, _ := regexp.Compile("Game (\\d+): ")
		idmatch := idRegex.FindStringSubmatch(text)
		id, _ := strconv.Atoi(idmatch[1])
		text = text[len(idmatch[0]):]

		blues := checkCubes("blue", 14, text)
		reds := checkCubes("red", 12, text)
		greens := checkCubes("green", 13, text)
		if blues && reds && greens {
			out += id
		}

	}
	fmt.Println(out)
}
