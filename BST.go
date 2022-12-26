package main

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

/**
 * @brief: Insert a new node into the BST.
 * @param root: The root of binary tree.
 * @param value: The value of the new node.
 * @return: The root of the new binary tree.
 */
func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{value: value}
	}

	if value < root.value {
		root.left = insert(root.left, value)
	} else {
		root.right = insert(root.right, value)
	}

	return root
}

/**
 * @brief: Find the node with the given value.
 * @param root: The root of binary tree.
 * @param value: The value of the node to be found.
 * @return: The node with the given value.
 */
func find(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	if root.value == value {
		return root
	}

	if value < root.value {
		return find(root.left, value)
	} else {
		return find(root.right, value)
	}
}
