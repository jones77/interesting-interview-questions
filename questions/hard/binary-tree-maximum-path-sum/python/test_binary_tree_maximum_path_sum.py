import unittest

from binary_tree_maximum_path_sum import TreeNode, max_path_sum


class TestMaxPathSum(unittest.TestCase):
    def test_examples(self):
        # Tree: [1, 2, 3]
        root1 = TreeNode(1)
        root1.left = TreeNode(2)
        root1.right = TreeNode(3)

        # Tree: [-10, 9, 20, null, null, 15, 7]
        root2 = TreeNode(-10)
        root2.left = TreeNode(9)
        root2.right = TreeNode(20)
        root2.right.left = TreeNode(15)
        root2.right.right = TreeNode(7)

        # Single node
        root3 = TreeNode(-3)

        tests = [
            (root1, 6),
            (root2, 42),
            (root3, -3),
        ]
        for root, want in tests:
            with self.subTest(root=root):
                self.assertEqual(max_path_sum(root), want)


if __name__ == "__main__":
    unittest.main()
