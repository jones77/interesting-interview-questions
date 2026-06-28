package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Question struct {
	ShortName    string `yaml:"short_name"`
	Title        string `yaml:"title"`
	QuestionText string `yaml:"question_text"`
	Hint         string `yaml:"hint,omitempty"`
	Difficulty   string `yaml:"difficulty"`
	Category     string `yaml:"category"`
	Solution     string `yaml:"-"`
}

func main() {
	questionsRaw, err := os.ReadFile("questions.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var questions []Question
	if err := yaml.Unmarshal(questionsRaw, &questions); err != nil {
		log.Fatal(err)
	}

	solutions := readSolutions()

	merged := 0
	for i, q := range questions {
		if sol, ok := solutions[q.ShortName]; ok {
			questions[i].Solution = sol
			merged++
		}
	}

	if merged != len(questions) {
		log.Printf("warning: merged %d solutions for %d questions", merged, len(questions))
	}

	for _, q := range questions {
		writeQuestionFile(q)
	}

	fmt.Printf("wrote %d question files\n", len(questions))
}

func readSolutions() map[string]string {
	raw, err := os.ReadFile("solutions.yaml")
	if err != nil {
		log.Fatal(err)
	}
	decoded := rot13(string(raw))
	var m map[string]string
	if err := yaml.Unmarshal([]byte(decoded), &m); err != nil {
		log.Fatal(err)
	}
	return m
}

func rot13(s string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case 'a' <= r && r <= 'z':
			return 'a' + (r-'a'+13)%26
		case 'A' <= r && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		default:
			return r
		}
	}, s)
}

func writeQuestionFile(q Question) {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("# %s\n\n", q.Title))
	b.WriteString(q.QuestionText)
	b.WriteString("\n")

	if q.Hint != "" {
		b.WriteString(fmt.Sprintf("\n> **Hint** — %s", q.Hint))
		b.WriteString("\n")
	}

	path := q.ShortName + ".md"
	if err := os.WriteFile(path, []byte(b.String()), 0644); err != nil {
		log.Fatalf("writing %s: %v", path, err)
	}
}
