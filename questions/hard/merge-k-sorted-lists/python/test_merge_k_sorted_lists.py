import unittest

from merge_k_sorted_lists import ListNode, merge_k_lists


def build_list(vals: list[int]) -> ListNode | None:
    if not vals:
        return None
    head = ListNode(vals[0])
    cur = head
    for val in vals[1:]:
        cur.next = ListNode(val)
        cur = cur.next
    return head


def list_to_slice(head: ListNode | None) -> list[int]:
    res = []
    while head:
        res.append(head.val)
        head = head.next
    return res


class TestMergeKLists(unittest.TestCase):
    def test_examples(self):
        tests = [
            (
                [
                    build_list([1, 4, 5]),
                    build_list([1, 3, 4]),
                    build_list([2, 6]),
                ],
                [1, 1, 2, 3, 4, 4, 5, 6],
            ),
            ([build_list([])], []),
            ([], []),
        ]
        for lists, want in tests:
            with self.subTest(lists=lists):
                self.assertEqual(list_to_slice(merge_k_lists(lists)), want)


if __name__ == "__main__":
    unittest.main()
