package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input")
	content := string(file)
	crates, commands, _ := strings.Cut(content, "\n\n")
	stacks := make([][]string, 9)
	crateRegex, _ := regexp.Compile("(\\[[A-Z]\\]| {4})")
	for _, line := range strings.Split(crates, "\n") {

		matches := crateRegex.FindAllString(line, -1)
		for i, match := range matches {
			if strings.Trim(match, " ") != "" {
				stacks[i] = append(stacks[i], strings.Trim(match, "[]"))
			}

		}
	}
	for _, stack := range stacks {
		fmt.Println(stack)
	}
	fmt.Println()
	cmdRegex, _ := regexp.Compile("move (\\d{1,2}) from (\\d{1,2}) to (\\d{1,2})")
	for _, cmd := range strings.Split(commands, "\n") {
		matches := cmdRegex.FindStringSubmatch(cmd)
		count, _ := strconv.Atoi(matches[1])
		start, _ := strconv.Atoi(matches[2])
		start--
		end, _ := strconv.Atoi(matches[3])
		end--
		for i := 0; i < count; i++ {
			stacks[end] = append([]string{stacks[start][0]}, stacks[end]...)

			stacks[start] = stacks[start][1:]
		}
	}
	output := ""
	for _, stack := range stacks {
		fmt.Println(stack)
		output += stack[0]
	}
	fmt.Println(output)

}
