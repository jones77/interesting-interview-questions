package main

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
	root1 := &TreeNode{Val: 2}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 3}

	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 4}
	root2.Right.Left = &TreeNode{Val: 3}
	root2.Right.Right = &TreeNode{Val: 6}

	root3 := &TreeNode{Val: 1}

	tests := []struct {
		root *TreeNode
		want bool
	}{
		{root1, true},
		{root2, false},
		{root3, true},
	}
	for _, tt := range tests {
		got := IsValidBST(tt.root)
		if got != tt.want {
			t.Errorf("IsValidBST(%v) = %t, want %t", tt.root, got, tt.want)
		}
	}
}
