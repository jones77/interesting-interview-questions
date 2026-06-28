package invertbinarytree_test

import (
	"github.com/jones/interesting-interview-questions/internal/types"
	"github.com/jones/interesting-interview-questions/questions/easy/invert-binary-tree"
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	root := &types.TreeNode{Val: 4}
	root.Left = &types.TreeNode{Val: 2}
	root.Right = &types.TreeNode{Val: 7}
	root.Left.Left = &types.TreeNode{Val: 1}
	root.Left.Right = &types.TreeNode{Val: 3}
	root.Right.Left = &types.TreeNode{Val: 6}
	root.Right.Right = &types.TreeNode{Val: 9}

	want := &types.TreeNode{Val: 4}
	want.Left = &types.TreeNode{Val: 7}
	want.Right = &types.TreeNode{Val: 2}
	want.Left.Left = &types.TreeNode{Val: 9}
	want.Left.Right = &types.TreeNode{Val: 6}
	want.Right.Left = &types.TreeNode{Val: 3}
	want.Right.Right = &types.TreeNode{Val: 1}

	got := invertbinarytree.InvertTree(root)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("invertbinarytree.InvertTree(%v) = %v, want %v", root, got, want)
	}
}
