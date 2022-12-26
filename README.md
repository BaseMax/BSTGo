# BST Go

This is a Go implementation of the BST data structure with a few of the most common operations. The algorithms code should be easy to understand.

## Structure

**Types:**

```go
type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}
```

**Functions:**

- `func insert(root *Node, value int) *Node`: Insert a new node into the BST.
- `func insertNonRecursively(root *Node, value int) *Node`: Insert a new node into the BST (non-recursively).
- `func find(root *Node, value int) *Node`: Find the node with the given value.
- `func findNonRecursively(root *Node, value int) *Node`: Find the node with the given value (non-recursively).
- `func findLeftMost(root *Node) *Node`: Find the left most child of the given node.
- `func findParent(root *Node, value int) *Node: Find the parent of the given node.
- `func findFirstParent(root *Node) *Node`: Find the first parent of the given node which is the left child of its parent.
- `func findNext(root *Node, value int) *Node`: Find next of the given node.
- `func findNextAny(root *Node, node *Node) *Node`: Find next of the given node (for any Binary Tree).
- `func deleteNode(root *Node, value int) *Node`: Delete a node with the given value.
- `func deleteNodeNonRecursively(root *Node, value int) *Node`: Delete a node with the given value (non-recursive).

- `func getHeight(root *Node) int`: Get the height of the BST.
- `func max(a, b int) int`: Get the maximum of two integers.

## Example

```go
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
```

Copyright (c) 2022, Max Base
