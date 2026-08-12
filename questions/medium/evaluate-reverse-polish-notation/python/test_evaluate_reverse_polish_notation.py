import unittest

from evaluate_reverse_polish_notation import eval_rpn


class TestEvalRPN(unittest.TestCase):
    def test_examples(self):
        tests = [
            (["2", "1", "+", "3", "*"], 9),
            (["4", "13", "5", "/", "+"], 6),
            (["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"], 22),
        ]
        for tokens, want in tests:
            with self.subTest(tokens=tokens):
                got = eval_rpn(tokens)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
