package main

import (
	"reflect"
	"testing"
)

func TestSerializeDeserialize(t *testing.T) {
	codec := Constructor()

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	serialized := codec.serialize(root)
	deserialized := codec.deserialize(serialized)

	if !reflect.DeepEqual(root, deserialized) {
		t.Errorf("round-trip failed:\n  original:    %v\n  deserialized: %v", root, deserialized)
	}

	empty := codec.serialize(nil)
	if got := codec.deserialize(empty); got != nil {
		t.Error("expected nil for empty tree")
	}
}
