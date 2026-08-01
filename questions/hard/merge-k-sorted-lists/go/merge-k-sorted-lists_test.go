package main

import (
	"reflect"
	"testing"
)

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{Val: vals[i]}
		cur = cur.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		lists []*ListNode
		want  []int
	}{
		{
			[]*ListNode{
				buildList([]int{1, 4, 5}),
				buildList([]int{1, 3, 4}),
				buildList([]int{2, 6}),
			},
			[]int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{[]*ListNode{buildList([]int{})}, nil},
		{[]*ListNode{}, nil},
	}
	for _, tt := range tests {
		got := listToSlice(MergeKLists(tt.lists))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("MergeKLists(%v) = %v, want %v", tt.lists, got, tt.want)
		}
	}
}
