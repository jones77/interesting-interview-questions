package binarytreemaximumpathsum_test

import (
	"github.com/jones/interesting-interview-questions/internal/types"
	"github.com/jones/interesting-interview-questions/questions/hard/binary-tree-maximum-path-sum"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	// Tree: [1,2,3]
	root1 := &types.TreeNode{Val: 1}
	root1.Left = &types.TreeNode{Val: 2}
	root1.Right = &types.TreeNode{Val: 3}

	// Tree: [-10,9,20,null,null,15,7]
	root2 := &types.TreeNode{Val: -10}
	root2.Left = &types.TreeNode{Val: 9}
	root2.Right = &types.TreeNode{Val: 20}
	root2.Right.Left = &types.TreeNode{Val: 15}
	root2.Right.Right = &types.TreeNode{Val: 7}

	// Single node
	root3 := &types.TreeNode{Val: -3}

	tests := []struct {
		root *types.TreeNode
		want int
	}{
		{root1, 6},
		{root2, 42},
		{root3, -3},
	}
	for _, tt := range tests {
		got := binarytreemaximumpathsum.MaxPathSum(tt.root)
		if got != tt.want {
			t.Errorf("binarytreemaximumpathsum.MaxPathSum(%v) = %d, want %d", tt.root, got, tt.want)
		}
	}
}
