package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
	all := flag.Bool("all", false, "show questions with their reference solutions")
	flag.Parse()

	questionsRaw, err := os.ReadFile("questions.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var questions []Question
	if err := yaml.Unmarshal(questionsRaw, &questions); err != nil {
		log.Fatal(err)
	}

	if !*all {
		for _, q := range questions {
			fmt.Printf("%s (%s, %s)\n", q.ShortName, q.Difficulty, q.Category)
		}
		return
	}

	solutionsRaw, err := os.ReadFile("solutions.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var solutions map[string]string
	if err := yaml.Unmarshal(solutionsRaw, &solutions); err != nil {
		log.Fatal(err)
	}

	for _, q := range questions {
		fmt.Printf("%s (%s, %s)\n", q.ShortName, q.Difficulty, q.Category)
		fmt.Println("---")
		if sol, ok := solutions[q.ShortName]; ok {
			fmt.Println(sol)
		} else {
			fmt.Println("(no solution available)")
		}
		fmt.Println()
	}
}
