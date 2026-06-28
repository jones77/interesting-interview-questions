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

	matched := 0
	for _, q := range questions {
		if _, ok := solutions[q.ShortName]; ok {
			matched++
		}
	}

	if matched != len(questions) {
		log.Printf("warning: matched %d solutions for %d questions", matched, len(questions))
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
	var m map[string]string
	if err := yaml.Unmarshal(raw, &m); err != nil {
		log.Fatal(err)
	}
	return m
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

	dir := "questions/" + q.Difficulty + "/" + q.ShortName
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("creating %s: %v", dir, err)
	}
	path := dir + "/" + q.ShortName + ".md"
	if err := os.WriteFile(path, []byte(b.String()), 0644); err != nil {
		log.Fatalf("writing %s: %v", path, err)
	}
}
