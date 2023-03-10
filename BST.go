/**
 * @Project: Go: Binary Search Tree
 * @Author: Max Base
 * @Date: 2022-12-26
 * @License: GPL-3.0
 * @Repository: https://github.com/BaseMax/BSTGo
 */

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
// B(logn)
// Worst case O(n-1)
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
 * @brief: Insert a new node into the BST (non-recursively).
 * @param root: The root of binary tree.
 * @param value: The value of the new node.
 * @return: The root of the new binary tree.
 */
// B(logn)
// worst case O(n-1)
func insertNonRecursively(root *Node, value int) *Node {
	if root == nil {
		return &Node{value: value}
	}

	n := root
	for n != nil {
		if value < n.value {
			if n.left == nil {
				n.left = &Node{value: value}
				break
			}
			n = n.left
		} else {
			if n.right == nil {
				n.right = &Node{value: value}
				break
			}
			n = n.right
		}
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
 * @brief: Find the node with the given value (non-recursively).
 * @param root: The root of binary tree.
 * @param value: The value of the node to be found.
 * @return: The node with the given value.
 */
func findNonRecursively(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	n := root
	for n != nil {
		if n.value == value {
			return n
		}

		if value < n.value {
			n = n.left
		} else {
			n = n.right
		}
	}

	return nil
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

/**
 * @brief: Get the height of the binary tree.
 * @param root: The root of binary tree.
 * @return: The height of the binary tree.
 */
func getHeight(root *Node) int {
	if root == nil {
		return 0
	}

	return max(getHeight(root.left), getHeight(root.right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * @brief: Get the height of the binary tree. (non-recursively)
 * @param root: The root of binary tree.
 * @return: The height of the binary tree.
 */
func getHeightNonRecursive(root *Node) int {
	if root == nil {
		return 0
	}

	q := []*Node{root}
	height := 0
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]
			if n.left != nil {
				q = append(q, n.left)
			}
			if n.right != nil {
				q = append(q, n.right)
			}
		}
		height++
	}

	return height
}

/**
 * @brief: Get the size of the binary tree.
 * @param root: The root of binary tree.
 * @return: The size of the binary tree.
 */
 func getSize(root *Node) int {
	if root == nil {
		return 0
	}

	return getSize(root.left) + getSize(root.right) + 1
}
	
/**
 * @brief: Get the size of the binary tree. (non-recursively)
 * @param root: The root of binary tree.
 * @return: The size of the binary tree.
 */
func getSizeNonRecursive(root *Node) int {
	if root == nil {
		return 0
	}

	q := []*Node{root}
	size := 0
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		size++
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}
	}

	return size
}

/**
 * @brief: Delete a node with the given value.
 * @param root: The root of binary tree.
 * @param value: The value of the node to be deleted.
 * @return: The root of the binary tree.
 */
func deleteNode(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	if value < root.value {
		root.left = deleteNode(root.left, value)
	}
	else if value > root.value {
		root.right = deleteNode(root.right, value)
	}
	// root.value is the value to be deleted
	else {
		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}

		// find the left most child of root.right
		n := findLeftMost(root.right)
		root.value = n.value
		root.right = deleteNode(root.right, n.value)
	}

	return root
}

/**
 * @brief: Delete a node with the given value (non-recursive).
 * @param root: The root of binary tree.
 * @param value: The value of the node to be deleted.
 * @return: The root of the binary tree.
 */
// Worst case: O(n)
// Average case: O(logn)
// Best case: O(1)
func deleteNodeNonRecursively(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	n := root
	var p *Node

	for n != nil && n.value != value {
		p = n
		if value < n.value {
			n = n.left
		} else {
			n = n.right
		}
	}

	// not found
	if n == nil {
		return root
	}
	// found and there is no left and right childs
	else if n.left == nil && n.right == nil {
		if p == nil {
			return nil
		}

		if p.left == n {
			p.left = nil
		} else {
			p.right = nil
		}
		return root
	}
	// found and there is no left child
	else if n.left == nil {
		if p == nil {
			return n.right
		}

		if p.left == n {
			p.left = n.right
		} else {
			p.right = n.right
		}
		return root
	}
	// found and there is no right child
	else if n.right == nil {
		if p == nil {
			return n.left
		}

		if p.left == n {
			p.left = n.left
		} else {
			p.right = n.left
		}
		return root
	}

	// found and there are both left and right child
	// find the left most child of n.right
	m := n.right
	var mp *Node
	for m.left != nil {
		mp = m
		m = m.left
	}

	n.value = m.value
	if mp == nil {
		n.right = m.right
	} else {
		mp.left = m.right
	}

	return root
}

func main() {
	// Create a BST
	bst := BST{}

	// Insert nodes into the BST
	bst.root = insert(bst.root, 5)
	bst.root = insert(bst.root, 3)
	bst.root = insert(bst.root, 7)
	bst.root = insert(bst.root, 2)
	bst.root = insert(bst.root, 4)
	bst.root = insert(bst.root, 6)
	bst.root = insert(bst.root, 8)

	// insertNonRecursively
	bst.root = insertNonRecursively(bst.root, 15)
	bst.root = insertNonRecursively(bst.root, 13)
	bst.root = insertNonRecursively(bst.root, 17)
	bst.root = insertNonRecursively(bst.root, 12)
	bst.root = insertNonRecursively(bst.root, 14)

	// Find the node with the given value
	n := findNext(bst.root, 5)
	if n != nil {
		println(n.value)
	}

	// Find a node
	n = find(bst.root, 5)
	if n != nil {
		println(n.value)
	}

	// findNonRecursively
	n = findNonRecursively(bst.root, 5)
	if n != nil {
		println(n.value)
	}

	// Get the height
	println(getHeight(bst.root))
}
