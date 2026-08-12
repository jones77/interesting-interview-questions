import unittest

from time_based_key_value_store import TimeMap


class TestTimeMap(unittest.TestCase):
    def test_example(self):
        tm = TimeMap()
        tm.set("foo", "bar", 1)
        self.assertEqual(tm.get("foo", 1), "bar")
        self.assertEqual(tm.get("foo", 3), "bar")
        tm.set("foo", "bar2", 4)
        self.assertEqual(tm.get("foo", 4), "bar2")
        self.assertEqual(tm.get("foo", 5), "bar2")


if __name__ == "__main__":
    unittest.main()
