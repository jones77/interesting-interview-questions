// Serialize and Deserialize Binary Tree
//
// Design an algorithm to serialize and deserialize a binary tree. There is no
// restriction on how your serialisation/deserialisation algorithm should
// work. You just need to ensure that a binary tree can be serialised to a
// string and this string can be deserialised back to the original tree.
//
// Example 1:
//   Input:  root = [1,2,3,null,null,4,5]
//   Output: [1,2,3,null,null,4,5]
// Example 2:
//   Input:  root = []
//   Output: []

package main

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
