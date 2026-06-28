package serializedeserializebinarytree_test

import (
	"github.com/jones/interesting-interview-questions/internal/types"
	"github.com/jones/interesting-interview-questions/questions/hard/serialize-deserialize-binary-tree"
	"reflect"
	"testing"
)

func TestSerializeDeserialize(t *testing.T) {
	codec := serializedeserializebinarytree.Constructor()

	root := &types.TreeNode{Val: 1}
	root.Left = &types.TreeNode{Val: 2}
	root.Right = &types.TreeNode{Val: 3}
	root.Right.Left = &types.TreeNode{Val: 4}
	root.Right.Right = &types.TreeNode{Val: 5}

	serialized := codec.Serialize(root)
	deserialized := codec.Deserialize(serialized)

	if !reflect.DeepEqual(root, deserialized) {
		t.Errorf("round-trip failed:\n  original:    %v\n  deserialized: %v", root, deserialized)
	}

	empty := codec.Serialize(nil)
	if got := codec.Deserialize(empty); got != nil {
		t.Error("expected nil for empty tree")
	}
}
