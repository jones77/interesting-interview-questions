package validatebinarysearchtree_test

import (
	"github.com/jones/interesting-interview-questions/internal/types"
	"github.com/jones/interesting-interview-questions/questions/medium/validate-binary-search-tree"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	root1 := &types.TreeNode{Val: 2}
	root1.Left = &types.TreeNode{Val: 1}
	root1.Right = &types.TreeNode{Val: 3}

	root2 := &types.TreeNode{Val: 5}
	root2.Left = &types.TreeNode{Val: 1}
	root2.Right = &types.TreeNode{Val: 4}
	root2.Right.Left = &types.TreeNode{Val: 3}
	root2.Right.Right = &types.TreeNode{Val: 6}

	root3 := &types.TreeNode{Val: 1}

	tests := []struct {
		root *types.TreeNode
		want bool
	}{
		{root1, true},
		{root2, false},
		{root3, true},
	}
	for _, tt := range tests {
		got := validatebinarysearchtree.IsValidBST(tt.root)
		if got != tt.want {
			t.Errorf("validatebinarysearchtree.IsValidBST(%v) = %t, want %t", tt.root, got, tt.want)
		}
	}
}
