import unittest

from invert_binary_tree import TreeNode, invert_tree


def _trees_equal(a: TreeNode | None, b: TreeNode | None) -> bool:
    if a is None or b is None:
        return a is b
    return a.val == b.val and _trees_equal(a.left, b.left) and _trees_equal(a.right, b.right)


class TestInvertTree(unittest.TestCase):
    def test_invert_tree(self):
        root = TreeNode(4)
        root.left = TreeNode(2)
        root.right = TreeNode(7)
        root.left.left = TreeNode(1)
        root.left.right = TreeNode(3)
        root.right.left = TreeNode(6)
        root.right.right = TreeNode(9)

        want = TreeNode(4)
        want.left = TreeNode(7)
        want.right = TreeNode(2)
        want.left.left = TreeNode(9)
        want.left.right = TreeNode(6)
        want.right.left = TreeNode(3)
        want.right.right = TreeNode(1)

        got = invert_tree(root)
        self.assertTrue(_trees_equal(got, want))


if __name__ == "__main__":
    unittest.main()
