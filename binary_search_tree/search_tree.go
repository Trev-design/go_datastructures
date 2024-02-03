package search_tree

import (
	"fmt"
)

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (node *Node) Insert(key int) error {
	if node.key < key {
		if node.right == nil {
			node.right = &Node{key: key}
			return nil
		} else {
			return node.right.Insert(key)
		}
	} else if node.key > key {
		if node.left == nil {
			node.left = &Node{key: key}
			return nil
		} else {
			return node.left.Insert(key)
		}
	}

	return fmt.Errorf("key already exist")
}

func (node *Node) Search(key int) bool {
	if node == nil {
		return false
	}

	if node.key < key {
		return node.right.Search(key)
	} else if node.key > key {
		return node.left.Search(key)
	}

	return true
}
