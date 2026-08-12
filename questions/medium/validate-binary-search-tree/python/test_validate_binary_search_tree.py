import unittest

from validate_binary_search_tree import TreeNode, is_valid_bst


class TestIsValidBST(unittest.TestCase):
    def test_examples(self):
        root1 = TreeNode(2)
        root1.left = TreeNode(1)
        root1.right = TreeNode(3)

        root2 = TreeNode(5)
        root2.left = TreeNode(1)
        root2.right = TreeNode(4)
        root2.right.left = TreeNode(3)
        root2.right.right = TreeNode(6)

        root3 = TreeNode(1)

        tests = [
            (root1, True),
            (root2, False),
            (root3, True),
        ]
        for root, want in tests:
            with self.subTest(root=root):
                got = is_valid_bst(root)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
