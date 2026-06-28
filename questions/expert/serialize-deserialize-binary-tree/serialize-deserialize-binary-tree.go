package main

// TODO: Implement TreeNode and Codec with Constructor, serialize, deserialize.
// See serialize-deserialize-binary-tree_test.go for expected behaviour.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	return ""
}

func (c *Codec) deserialize(data string) *TreeNode {
	return nil
}

func main() {}
