package engine

import (
	"fmt"
	"strings"
)

type node struct {
	pattern    string
	searchPath string
	children   []*node // Children array
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, searchPath=%s}", n.pattern, n.searchPath)
}

// Insert node recursion, assign pattern to the leaf node' while branch node keep
// an empty pattern
func (n *node) insert(pattern string, searchPath []string, height int) bool {
	// fmt.Println("pattern insert = %s, searchPath = %s, height = %d \n", pattern, searchPath, height)
	if len(searchPath) == height {
		fmt.Println("child is =", n)
		// fmt.Println("n.pattern = %s", n.pattern)
		n.pattern = pattern
		return true
	}

	item := searchPath[height]
	child := n.matchChild(item)
	if child == nil {
		child = &node{
			pattern: pattern, searchPath: item}
		n.children = append(n.children, child)
	}
	return child.insert(pattern, searchPath, height+1)
}

// Search node recursion
func (n *node) search(pattern string, searchPath []string, height int) *node {
	if len(searchPath) == height || strings.HasPrefix(n.searchPath, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	item := searchPath[height]
	children := n.matchChildren(item)

	for _, child := range children {
		res := child.search(pattern, searchPath, height+1)
		if res != nil {
			return res
		}
	}
	return nil
}

// Return node who have search path firstly
func (n *node) matchChild(searchPath string) *node {
	for _, child := range n.children {
		if child.pattern == searchPath {
			return child
		}
	}
	return nil
}

// Return nodes who have search path
func (n *node) matchChildren(searchPath string) []*node {
	returnNodes := make([]*node, 0)
	for _, child := range n.children {
		if child.pattern == searchPath {
			returnNodes = append(returnNodes, child)
		}
	}
	return returnNodes
}
