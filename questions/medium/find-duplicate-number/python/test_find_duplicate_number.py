import unittest

from find_duplicate_number import find_duplicate


class TestFindDuplicate(unittest.TestCase):
    def test_examples(self):
        tests = [
            ([1, 3, 4, 2, 2], 2),
            ([3, 1, 3, 4, 2], 3),
            ([1, 1], 1),
        ]
        for nums, want in tests:
            with self.subTest(nums=nums):
                got = find_duplicate(nums)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
