package router

import (
	"fmt"
	"strings"
)

const (
	ITEMS = 39
)

type (
	Node struct {
		Children [ITEMS]*Node
		hasParam bool
		Param    string
		Value    interface{}
		End      bool
	}
)

func NewTree() *Node {
	return &Node{
		Children: [ITEMS]*Node{},
	}
}

func newNode() *Node {
	return &Node{
		Children: [ITEMS]*Node{},
	}
}

func (n *Node) Insert(s string, v interface{}) {
	node := n

	hasParam := false

	for _, c := range strings.ToLower(s) {
		// If we have a param, we need to check if the next character is a slash
		if !hasParam || c == '/' {
			var idx int
			switch {
			case c == '/':
				idx = 36
				hasParam = false
			case c == ':':
				idx = 37
				hasParam = true
			case c == '*':
				idx = 38
			case c >= 'a' && c <= 'z':
				idx = int(c) - 'a'
			case c >= '0' && c <= '9':
				idx = int(c) - '0' + 26
			default:
				panic(fmt.Sprintf("Invalid character: %v\n", c))
			}

			if node.Children[idx] == nil {
				node.Children[idx] = newNode()
			}

			node = node.Children[idx]
			continue
		} else {
			// If we have a param character, we need to add it to the param string
			node.Param += string(c)
			node.hasParam = true
		}

	}

	node.Value = v
	node.End = true

}

func (n *Node) Search(s string) (interface{}, map[string]string) {
	node := n

	params := make(map[string]string)

	for _, c := range strings.ToLower(s) {
		if !node.hasParam || c == '/' {
			var idx int
			switch {
			case c == '/':
				idx = 36
			case c >= 'a' && c <= 'z':
				idx = int(c) - 'a'
			case c >= '0' && c <= '9':
				idx = int(c) - '0' + 26
			default:
				panic(fmt.Sprintf("Invalid character: %v\n", c))
			}

			if node.Children[idx] == nil {
				return nil, nil
			}

			node = node.Children[idx]
			if node.Children[37] != nil {
				node = node.Children[37]
			}
		} else {
			// If we have a param character, we need to add it to the param string
			params[node.Param] += string(c)

		}

	}

	if node.End {
		return node.Value, params
	}

	return nil, nil
}
