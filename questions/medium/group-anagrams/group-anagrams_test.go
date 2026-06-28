package groupanagrams_test

import (
	"github.com/jones/interesting-interview-questions/questions/medium/group-anagrams"
	"reflect"
	"sort"
	"testing"
)

func groupsEqual(got, want [][]string) bool {
	if len(got) != len(want) {
		return false
	}
	normalize := func(g [][]string) [][]string {
		for i := range g {
			sort.Strings(g[i])
		}
		sort.Slice(g, func(i, j int) bool {
			for k := 0; k < len(g[i]) && k < len(g[j]); k++ {
				if g[i][k] != g[j][k] {
					return g[i][k] < g[j][k]
				}
			}
			return len(g[i]) < len(g[j])
		})
		return g
	}
	gotCopy := make([][]string, len(got))
	for i := range got {
		gotCopy[i] = append([]string{}, got[i]...)
	}
	wantCopy := make([][]string, len(want))
	for i := range want {
		wantCopy[i] = append([]string{}, want[i]...)
	}
	return reflect.DeepEqual(normalize(gotCopy), normalize(wantCopy))
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		got := groupanagrams.GroupAnagrams(tt.strs)
		if !groupsEqual(got, tt.want) {
			t.Errorf("groupanagrams.GroupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
		}
	}
}
