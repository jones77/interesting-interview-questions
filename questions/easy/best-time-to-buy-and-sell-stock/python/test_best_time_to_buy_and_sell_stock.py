import unittest

from best_time_to_buy_and_sell_stock import max_profit


class TestMaxProfit(unittest.TestCase):
    def test_max_profit(self):
        tests = [
            ([7, 1, 5, 3, 6, 4], 5),
            ([7, 6, 4, 3, 1], 0),
            ([1], 0),
            ([1, 2], 1),
        ]
        for prices, want in tests:
            with self.subTest(prices=prices):
                self.assertEqual(max_profit(prices), want)


if __name__ == "__main__":
    unittest.main()
