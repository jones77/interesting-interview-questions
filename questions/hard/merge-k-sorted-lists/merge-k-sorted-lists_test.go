package mergeksortedlists_test

import (
	"github.com/jones/interesting-interview-questions/internal/types"
	"github.com/jones/interesting-interview-questions/questions/hard/merge-k-sorted-lists"
	"reflect"
	"testing"
)

func buildList(vals []int) *types.ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &types.ListNode{Val: vals[0]}
	cur := head
	for i := 1; i < len(vals); i++ {
		cur.Next = &types.ListNode{Val: vals[i]}
		cur = cur.Next
	}
	return head
}

func listToSlice(head *types.ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		lists []*types.ListNode
		want  []int
	}{
		{
			[]*types.ListNode{
				buildList([]int{1, 4, 5}),
				buildList([]int{1, 3, 4}),
				buildList([]int{2, 6}),
			},
			[]int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{[]*types.ListNode{buildList([]int{})}, nil},
		{[]*types.ListNode{}, nil},
	}
	for _, tt := range tests {
		got := listToSlice(mergeksortedlists.MergeKLists(tt.lists))
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("mergeksortedlists.MergeKLists(%v) = %v, want %v", tt.lists, got, tt.want)
		}
	}
}
