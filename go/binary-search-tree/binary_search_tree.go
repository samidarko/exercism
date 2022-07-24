package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	switch {
	case i <= bst.data && bst.left == nil:
		bst.left = NewBst(i)
	case i > bst.data && bst.right == nil:
		bst.right = NewBst(i)
	case i <= bst.data && bst.left != nil:
		bst.left.Insert(i)
	case i > bst.data && bst.right != nil:
		bst.right.Insert(i)
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	result := make([]int, 0)
	inOrderTraversal(bst, &result)
	return result
}

func inOrderTraversal(bst *BinarySearchTree, result *[]int) {
	if bst != nil {
		inOrderTraversal(bst.left, result)
		*result = append(*result, bst.data)
		inOrderTraversal(bst.right, result)
	}
}
