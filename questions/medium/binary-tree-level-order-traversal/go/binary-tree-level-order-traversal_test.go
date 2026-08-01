package main

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}

	root2 := &TreeNode{Val: 1}

	tests := []struct {
		root *TreeNode
		want [][]int
	}{
		{root1, [][]int{{3}, {9, 20}, {15, 7}}},
		{nil, nil},
		{root2, [][]int{{1}}},
	}
	for _, tt := range tests {
		got := LevelOrder(tt.root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("LevelOrder(%v) = %v, want %v", tt.root, got, tt.want)
		}
	}
}
