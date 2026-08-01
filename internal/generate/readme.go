package generate

import (
	"strings"

	"github.com/jones/interesting-interview-questions/internal/question"
)

func renderREADME(s question.Spec) string {
	var b strings.Builder
	b.WriteString("# ")
	b.WriteString(s.Title)
	b.WriteString("\n")
	if s.URL != "" {
		b.WriteString("\nSource: ")
		b.WriteString(s.URL)
		b.WriteString("\n")
	}
	b.WriteString("\n")
	b.WriteString(strings.TrimSpace(s.QuestionText))
	b.WriteString("\n")
	return b.String()
}
