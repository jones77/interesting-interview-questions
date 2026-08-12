import unittest

from binary_tree_level_order_traversal import TreeNode, level_order


class TestLevelOrder(unittest.TestCase):
    def test_examples(self):
        root1 = TreeNode(3)
        root1.left = TreeNode(9)
        root1.right = TreeNode(20)
        root1.right.left = TreeNode(15)
        root1.right.right = TreeNode(7)

        root2 = TreeNode(1)

        tests = [
            (root1, [[3], [9, 20], [15, 7]]),
            (None, []),
            (root2, [[1]]),
        ]
        for root, want in tests:
            with self.subTest(root=root):
                got = level_order(root)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
