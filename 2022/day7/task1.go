package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"rutmanz/aoc/tree"
)

var root *tree.TreeNode = &tree.TreeNode{
	Name:   "/",
	Size:   0,
	Parent: nil,
}

var currentNode *tree.TreeNode = root

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text[0] == '$' {
			processCommand(text)
		} else {
			processChild(text)
		}
	}
	// fmt.Println(root)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	totalSize := root.GetSize()
	best := totalSize
	for _, dir := range root.GetChildDirs() {
		size := dir.GetSize()

		if size < best && totalSize-size <= 40000000 && dir.IsDir() {
			fmt.Println(dir, totalSize-size, best)
			best = size
		}
	}
	fmt.Println("!!!", best)

}
func processChild(text string) {
	split := strings.Split(text, " ")
	size := 0
	if split[0] != "dir" {
		size, _ = strconv.Atoi(split[0])
	}
	currentNode.AddChild(tree.TreeNode{
		Name:   split[1],
		Size:   size,
		Parent: currentNode,
	})
}
func processCommand(text string) {
	cmd := text[2:]

	if cmd[0:2] == "cd" {
		currentNode = getDir(cmd[3:])
		fmt.Println(currentNode)
	}

}

func getDir(name string) *tree.TreeNode {
	if name == "/" {
		return root
	} else if name == ".." {
		return currentNode.Parent
	} else {
		return currentNode.GetChild(name)
	}
}
