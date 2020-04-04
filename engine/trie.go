package engine

import "strings"

type node struct {
	pattern string
	searchPath string
	children []*node   // Children array 
}

// Insert node recursion, assign pattern to the leaf node' while branch node keep 
// an empty pattern
func (n *node) insert(pattern string, search_path []string, height int) {
	if len(search_path) == height {
		n.pattern = pattern
		return
	}

	item := search_path[height]
	child := n.matchChild(item)
	if child == nil {
		node := node { 
			searchPath: item}
		n.children = append(n.children, &node)
	}
	child.insert(pattern, search_path, height + 1)
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