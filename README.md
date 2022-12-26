# BST Go

This is a Go implementation of the BST data structure with a few of the most common operations.

## Structure


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
