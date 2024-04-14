package gom

import (
	"fmt"
	"strings"
)

type node struct {
	// pattern routes to be matched
	// e.g. /p/:lang
	pattern string
	// part of the route
	// e.g. :lang
	part string
	// child node in the trie tree
	// e.g. [doc, tutorial, intro]
	children []*node
	// Whether it's a precise match
	// true when part contains : or *
	isWild bool
}

// String print the node info
func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t", n.pattern, n.part, n.isWild)
}

// The first matching node, used for insertion
// part represents the matched route
// the return value is the first node pointer that satisfies the rule
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

// All successfully matched nodes, used for searching
// part represents the matched route
// the return value is the node pointer slice that satisfies the rule
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

// travel the node pointer list
func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

// Register routes
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)

	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

// Search routes
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}

		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
