package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "modernc.org/sqlite"
)

type question struct {
	shortName    string
	title        string
	questionText string
	hint         sql.NullString
	difficulty   string
	category     string
}

func main() {
	db, err := sql.Open("sqlite", "interview-questions.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT short_name, title, question_text, hint, difficulty, category
		FROM questions
		ORDER BY category, short_name
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var questions []question
	for rows.Next() {
		var q question
		if err := rows.Scan(&q.shortName, &q.title, &q.questionText, &q.hint, &q.difficulty, &q.category); err != nil {
			log.Fatal(err)
		}
		questions = append(questions, q)
	}

	for _, q := range questions {
		writeQuestionFile(q)
	}

	fmt.Printf("wrote %d question files\n", len(questions))
}

func writeQuestionFile(q question) {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("# %s\n\n", q.title))
	b.WriteString(q.questionText)
	b.WriteString("\n")

	if q.hint.Valid && q.hint.String != "" {
		b.WriteString(fmt.Sprintf("\n> **Hint** — %s", q.hint.String))
		b.WriteString("\n")
	}

	path := q.shortName + ".md"
	if err := os.WriteFile(path, []byte(b.String()), 0644); err != nil {
		log.Fatalf("writing %s: %v", path, err)
	}
}
