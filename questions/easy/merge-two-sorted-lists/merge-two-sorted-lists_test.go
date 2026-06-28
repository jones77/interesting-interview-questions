package mergetwosortedlists_test

import (
	"github.com/jones/interesting-interview-questions/internal/types"
	"github.com/jones/interesting-interview-questions/questions/easy/merge-two-sorted-lists"
	"reflect"
	"testing"
)

func sliceToList(s []int) *types.ListNode {
	dummy := &types.ListNode{}
	cur := dummy
	for _, v := range s {
		cur.Next = &types.ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(l *types.ListNode) []int {
	var s []int
	for cur := l; cur != nil; cur = cur.Next {
		s = append(s, cur.Val)
	}
	return s
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1 []int
		list2 []int
		want  []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{nil, []int{1}, []int{1}},
		{nil, nil, nil},
	}
	for _, tt := range tests {
		l1 := sliceToList(tt.list1)
		l2 := sliceToList(tt.list2)
		got := listToSlice(mergetwosortedlists.MergeTwoLists(l1, l2))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("mergetwosortedlists.MergeTwoLists(%v, %v) = %v, want %v", tt.list1, tt.list2, got, tt.want)
		}
	}
}
