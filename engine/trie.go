package engine

import (
	"fmt"
	"strings"
)

type node struct {
	pattern string
	searchPath string
	children []*node   // Children array 
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, searchPath=%s}", n.pattern, n.searchPath)
}

// Insert node recursion, assign pattern to the leaf node' while branch node keep 
// an empty pattern
func (n *node) insert(pattern string, search_path []string, height int) bool {
	fmt.Println("pattern insert = %s, search_path = %s, height = %d \n", pattern, search_path, height)
	if len(search_path) == height {
		fmt.Println("child is =", n)
		// fmt.Println("n.pattern = %s", n.pattern)
		n.pattern = pattern
		return true
	}

	item := search_path[height]
	child := n.matchChild(item)
	if child == nil {
		child = &node { 
			searchPath: item}
		n.children = append(n.children, child)
	}
	return child.insert(pattern, search_path, height + 1)
}

// Search node recursion 
func (n *node) search(pattern string, search_path []string, height int) *node {
	if len(search_path) == height || strings.HasPrefix(n.searchPath, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	item := search_path[height]
	children := n.matchChildren(item)

	for _, child := range children {
		res := child.search(pattern, search_path, height + 1)
		if res != nil {
			return res
		}
	}
	return nil
}

// Return node who have search path firstly
func (n *node) matchChild(search_path string) *node {
	for _, child := range n.children {
		if child.pattern == search_path {
			return child 
		}
	}
	return nil
}

// Return nodes who have search path
func (n *node) matchChildren(search_path string) []*node {
	res_nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.pattern == search_path {
			res_nodes = append(res_nodes, child)
		}
	}
	return res_nodes
}