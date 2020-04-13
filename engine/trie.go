package engine

import (
	"fmt"
	"strings"
)

type node struct {
	pattern    string
	searchPath string
	children   []*node // Children array
	isWild     bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, searchPath=%s, isWild=%t}", n.pattern, n.searchPath, n.isWild)
}

// Insert node recursion, assign pattern to the leaf node' while branch node keep
// an empty pattern
func (n *node) insert(pattern string, searchPath []string, height int) bool {
	// fmt.Println("pattern insert = %s, searchPath = %s, height = %d \n", pattern, searchPath, height)
	if len(searchPath) == height {
		// fmt.Println("n.pattern = %s", n.pattern)
		n.pattern = pattern
		fmt.Println("child is =", n)
		return true
	}

	item := searchPath[height]
	child := n.matchChild(item)
	if child == nil {
		child = &node{
			searchPath: item, isWild: item[0] == ':' || item[0] == '*'}
		n.children = append(n.children, child)
	}
	return child.insert(pattern, searchPath, height+1)
}

// Search node recursion
func (n *node) search(searchPath []string, height int) *node {
	fmt.Println("**** searchPath: ", searchPath, "height: ", height)
	if len(searchPath) == height || strings.HasPrefix(n.searchPath, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	item := searchPath[height]
	fmt.Println("**** n.children: ", n.children, "item: ", item)
	children := n.matchChildren(item)
	fmt.Println("**** children: ", children)

	for _, child := range children {
		res := child.search(searchPath, height+1)
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
		if child.searchPath == searchPath || child.isWild {
			returnNodes = append(returnNodes, child)
		}
	}
	return returnNodes
}
