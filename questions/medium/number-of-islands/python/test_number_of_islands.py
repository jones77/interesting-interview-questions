import unittest

from number_of_islands import num_islands


class TestNumIslands(unittest.TestCase):
    def test_examples(self):
        tests = [
            (
                [
                    ["1", "1", "1", "1", "0"],
                    ["1", "1", "0", "1", "0"],
                    ["1", "1", "0", "0", "0"],
                    ["0", "0", "0", "0", "0"],
                ],
                1,
            ),
            (
                [
                    ["1", "1", "0", "0", "0"],
                    ["1", "1", "0", "0", "0"],
                    ["0", "0", "1", "0", "0"],
                    ["0", "0", "0", "1", "1"],
                ],
                3,
            ),
            ([["0"]], 0),
        ]
        for grid, want in tests:
            with self.subTest(grid=grid):
                got = num_islands(grid)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
