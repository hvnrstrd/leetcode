package main

import "fmt"

// Maximum Depth of Binary Tree
// Complexity - Time: O(n) - Space: O(h) where h is tree height
// Pattern: DFS, Recursion
// Recursively find max depth of left and right subtrees, return max + 1.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// ------------------------------------------------------------------

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println(maxDepth(root))
}
