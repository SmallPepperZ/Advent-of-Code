package main

import (
	"fmt"
	"io/ioutil"
)

const length = 14

func main() {
	file, _ := ioutil.ReadFile("input")
	content := string(file)
	for i := range content {
		unique := true
		for i1, char1 := range content[i : i+length] {
			for _, char2 := range content[i : i+i1] {
				if char1 == char2 {
					unique = false
				}
			}
		}
		if unique {
			fmt.Println(i + length)
			return
		}
	}
}
