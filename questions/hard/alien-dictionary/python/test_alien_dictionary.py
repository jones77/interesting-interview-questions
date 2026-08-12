import unittest

from alien_dictionary import alien_order


class TestAlienOrder(unittest.TestCase):
    def test_examples(self):
        tests = [
            (["wrt", "wrf", "er", "ett", "rftt"], "wertf"),
            (["z", "x", "z"], ""),
            (["z", "x"], "zx"),
        ]
        for words, want in tests:
            with self.subTest(words=words):
                self.assertEqual(alien_order(words), want)


if __name__ == "__main__":
    unittest.main()
