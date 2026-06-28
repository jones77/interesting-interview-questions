package wordladder_test

import (
	"github.com/jones/interesting-interview-questions/questions/hard/word-ladder"
	"testing"
)

func TestLadderLength(t *testing.T) {
	tests := []struct {
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
		{"a", "c", []string{"a", "b", "c"}, 2},
	}
	for _, tt := range tests {
		got := wordladder.LadderLength(tt.beginWord, tt.endWord, tt.wordList)
		if got != tt.want {
			t.Errorf("wordladder.LadderLength(%q, %q, %v) = %d, want %d", tt.beginWord, tt.endWord, tt.wordList, got, tt.want)
		}
	}
}
