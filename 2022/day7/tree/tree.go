package tree

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Name     string
	Size     int
	Parent   *TreeNode
	children []*TreeNode
}

func (node *TreeNode) GetChild(name string) *TreeNode {
	for _, child := range node.children {
		if child.Name == name {
			return child
		}
	}
	return nil
}

func (node *TreeNode) AddChild(child TreeNode) {
	if node.GetChild(child.Name) == nil {
		node.children = append(node.children, &child)
	}
}

func (node *TreeNode) String() string {
	lines := []string{fmt.Sprintf("%v %s", node.GetSize(), node.Name)}
	for _, child := range node.children {
		lines = append(lines, "    -"+child.Parent.Name+"    "+child.String())
	}
	return strings.Join(lines, "\n")
}

func (node *TreeNode) GetSize() int {
	size := node.Size
	for _, child := range node.children {
		size += child.GetSize()
	}
	return size
}

func (node *TreeNode) GetChildDirs() []*TreeNode {
	nodes := node.children
	for _, child := range node.children {
		nodes = append(nodes, child.GetChildDirs()...)
	}
	return nodes
}

func (node *TreeNode) IsDir() bool {
	return node.Size == 0
}
