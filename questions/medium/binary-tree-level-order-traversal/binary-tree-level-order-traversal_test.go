package binarytreelevelordertraversal_test

import (
	"github.com/jones/interesting-interview-questions/internal/types"
	"github.com/jones/interesting-interview-questions/questions/medium/binary-tree-level-order-traversal"
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root1 := &types.TreeNode{Val: 3}
	root1.Left = &types.TreeNode{Val: 9}
	root1.Right = &types.TreeNode{Val: 20}
	root1.Right.Left = &types.TreeNode{Val: 15}
	root1.Right.Right = &types.TreeNode{Val: 7}

	root2 := &types.TreeNode{Val: 1}

	tests := []struct {
		root *types.TreeNode
		want [][]int
	}{
		{root1, [][]int{{3}, {9, 20}, {15, 7}}},
		{nil, nil},
		{root2, [][]int{{1}}},
	}
	for _, tt := range tests {
		got := binarytreelevelordertraversal.LevelOrder(tt.root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("binarytreelevelordertraversal.LevelOrder(%v) = %v, want %v", tt.root, got, tt.want)
		}
	}
}
