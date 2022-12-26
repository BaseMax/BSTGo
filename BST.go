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
// O(logn) (worst case O(n))
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

/**
 * @berif: Find next of the given node.
 * @param root: The root of binary tree.
 * @return: The next node of the given node.
 */
func findNext(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	node := find(root, value)
	if node == nil {
		return nil
	}

	// if the node has right child, find the leftmost node of the right child
	if node.right != nil {
		q := []*Node{node.right}
		for len(q) > 0 {
			n := q[0]
			q = q[1:]
			if n.left == nil {
				return n
			}
			q = append(q, n.left)
		}
	}

	// if the node has no right child, find the first node that is the left child of its parent
	p := root
	n := node
	for p != nil {
		if p.left == n {
			return p
		}
		if n.value < p.value {
			p = p.left
		} else {
			p = p.right
		}
	}
	return nil

/**
 * @berif: Find next of the given node (for any Binary Tree)
 * @param root: The root of binary tree.
 * @param node: The given node.
 * @return: The next node of the given node.
 */
func findNextAny(root *Node, node *Node) *Node {
	if root == nil {
		return nil
	}

	// for any binary tree, not only BST
	if root.right == nil {
		return nil
	}
	q := []*Node{root.right}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n.left == nil {
			return n
		}
		q = append(q, n.left)
	}
	return nil
}

/**
 * @berif: Insert a new node into the BST. (Sorted BST)
 * @param root: The root of binary tree.
 * @param value: The value of the new node.
 * @return: The root of the new binary tree.
 */
func insertSorted(root *Node, value int) *Node {
	if root == nil {
		return &Node{value: value}
	}

	if value < root.value {
		root.left = insertSorted(root.left, value)
	} else {
		root.right = insertSorted(root.right, value)
	}

	return root
}
