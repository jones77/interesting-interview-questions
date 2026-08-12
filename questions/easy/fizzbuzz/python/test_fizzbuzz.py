import unittest

from fizzbuzz import fizz_buzz


class TestFizzBuzz(unittest.TestCase):
    def test_fizz_buzz(self):
        tests = [
            (3, ["1", "2", "Fizz"]),
            (5, ["1", "2", "Fizz", "4", "Buzz"]),
            (
                15,
                [
                    "1",
                    "2",
                    "Fizz",
                    "4",
                    "Buzz",
                    "Fizz",
                    "7",
                    "8",
                    "Fizz",
                    "Buzz",
                    "11",
                    "Fizz",
                    "13",
                    "14",
                    "FizzBuzz",
                ],
            ),
        ]
        for n, want in tests:
            with self.subTest(n=n):
                self.assertEqual(fizz_buzz(n), want)


if __name__ == "__main__":
    unittest.main()
