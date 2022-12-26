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
// B(logn)
// worst case O(n-1)
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
 * @brief: Find the left most child of the given node.
 * @param root: The root of binary tree.
 * @return: The left most child of the given node.
 */
func findLeftMost(root *Node) *Node {
	if root == nil {
		return nil
	}

	for root.left != nil {
		root = root.left
	}

	return root
}

/**
 * @brief: Find the parent of the given node.
 * @param root: The root of binary tree.
 * @param value: The value of the given node.
 * @return: The parent of the given node.
 */
func findParent(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	if root.left != nil && root.left.value == value {
		return root
	}

	if root.right != nil && root.right.value == value {
		return root
	}

	if value < root.value {
		return findParent(root.left, value)
	} else {
		return findParent(root.right, value)
	}
}

/**
 * @brief: Find the first parent of the given node which is the left child of its parent.
 * @param root: The root of binary tree.
 * @return: The first parent of the given node which is the left child of its parent.
 */
func findFirstParent(root *Node) *Node {
	if root == nil {
		return nil
	}

	p := findParent(root, root.value)
	if p == nil {
		return nil
	}

	if p.left != nil && p.left.value == root.value {
		return p
	}

	return findFirstParent(p)
}

/**
 * @berif: Find next of the given node.
 * @param root: The root of binary tree.
 * @return: The next node of the given node.
 */
func findNext(root *Node, value int) *Node {
	n := find(root, value)

	if n == nil {
		return nil
	}

	// if n has right child, find the left most child of n.right
	if n.right != nil {
		return findLeftMost(n.right)
	}

	// if n has no right child, find the first parent of n which is the left child of its parent
	return findFirstParent(n)
}

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
