package main

import (
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	// Tree: [1,2,3]
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}

	// Tree: [-10,9,20,null,null,15,7]
	root2 := &TreeNode{Val: -10}
	root2.Left = &TreeNode{Val: 9}
	root2.Right = &TreeNode{Val: 20}
	root2.Right.Left = &TreeNode{Val: 15}
	root2.Right.Right = &TreeNode{Val: 7}

	// Single node
	root3 := &TreeNode{Val: -3}

	tests := []struct {
		root *TreeNode
		want int
	}{
		{root1, 6},
		{root2, 42},
		{root3, -3},
	}
	for _, tt := range tests {
		got := MaxPathSum(tt.root)
		if got != tt.want {
			t.Errorf("MaxPathSum(%v) = %d, want %d", tt.root, got, tt.want)
		}
	}
}
