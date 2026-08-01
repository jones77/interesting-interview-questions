package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec solves Serialize and Deserialize Binary Tree. See README.md for the problem statement.
type Codec struct{}

func Constructor() *Codec {
	return &Codec{}
}

func (c *Codec) Serialize(root *TreeNode) string {
	return ""
}

func (c *Codec) Deserialize(data string) *TreeNode {
	return nil
}

func main() {}
