import unittest

from serialize_deserialize_binary_tree import Codec, TreeNode


def tree_equal(a: TreeNode | None, b: TreeNode | None) -> bool:
    if a is None or b is None:
        return a is b
    return (
        a.val == b.val
        and tree_equal(a.left, b.left)
        and tree_equal(a.right, b.right)
    )


class TestSerializeDeserialize(unittest.TestCase):
    def test_round_trip(self):
        codec = Codec()

        root = TreeNode(1)
        root.left = TreeNode(2)
        root.right = TreeNode(3)
        root.right.left = TreeNode(4)
        root.right.right = TreeNode(5)

        serialized = codec.serialize(root)
        deserialized = codec.deserialize(serialized)

        self.assertTrue(tree_equal(root, deserialized))

    def test_empty_tree(self):
        codec = Codec()
        empty = codec.serialize(None)
        self.assertIsNone(codec.deserialize(empty))


if __name__ == "__main__":
    unittest.main()
