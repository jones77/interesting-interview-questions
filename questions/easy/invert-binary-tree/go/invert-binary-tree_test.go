package main

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}

	want := &TreeNode{Val: 4}
	want.Left = &TreeNode{Val: 7}
	want.Right = &TreeNode{Val: 2}
	want.Left.Left = &TreeNode{Val: 9}
	want.Left.Right = &TreeNode{Val: 6}
	want.Right.Left = &TreeNode{Val: 3}
	want.Right.Right = &TreeNode{Val: 1}

	got := InvertTree(root)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InvertTree(%v) = %v, want %v", root, got, want)
	}
}
