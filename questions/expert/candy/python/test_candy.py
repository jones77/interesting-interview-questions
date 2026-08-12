import unittest

from candy import candy


class TestCandy(unittest.TestCase):
    def test_candy(self):
        tests = [
            ([1, 0, 2], 5),
            ([1, 2, 2], 4),
            ([1], 1),
            ([1, 3, 2, 2, 1], 7),
        ]
        for ratings, want in tests:
            with self.subTest(ratings=ratings):
                got = candy(ratings)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
