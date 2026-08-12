import unittest

from merge_two_sorted_lists import ListNode, merge_two_lists


def _slice_to_list(values: list[int]) -> ListNode | None:
    dummy = ListNode()
    cur = dummy
    for v in values:
        cur.next = ListNode(v)
        cur = cur.next
    return dummy.next


def _list_to_slice(head: ListNode | None) -> list[int]:
    result = []
    cur = head
    while cur:
        result.append(cur.val)
        cur = cur.next
    return result


class TestMergeTwoLists(unittest.TestCase):
    def test_merge_two_lists(self):
        tests = [
            ([1, 2, 4], [1, 3, 4], [1, 1, 2, 3, 4, 4]),
            ([], [1], [1]),
            ([], [], []),
        ]
        for list1, list2, want in tests:
            with self.subTest(list1=list1, list2=list2):
                l1 = _slice_to_list(list1)
                l2 = _slice_to_list(list2)
                got = _list_to_slice(merge_two_lists(l1, l2))
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
