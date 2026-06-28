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
	var showQuestion bool
	flag.BoolVar(&showQuestion, "q", false, "include question text")
	flag.BoolVar(&showQuestion, "questions", false, "include question text")
	var showSolution bool
	flag.BoolVar(&showSolution, "s", false, "include reference solution")
	flag.BoolVar(&showSolution, "solutions", false, "include reference solution")
	flag.Parse()

	questionsRaw, err := os.ReadFile("questions.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var questions []Question
	if err := yaml.Unmarshal(questionsRaw, &questions); err != nil {
		log.Fatal(err)
	}

	var solutions map[string]string
	if showSolution {
		solutionsRaw, err := os.ReadFile("solutions.yaml")
		if err != nil {
			log.Fatal(err)
		}
		if err := yaml.Unmarshal(solutionsRaw, &solutions); err != nil {
			log.Fatal(err)
		}
	}

	for _, q := range questions {
		fmt.Printf("%s (%s, %s) — %s\n", q.ShortName, q.Difficulty, q.Category, q.Title)

		if showQuestion {
			fmt.Println()
			fmt.Println(q.QuestionText)
		}

		if showSolution {
			fmt.Println("---")
			if sol, ok := solutions[q.ShortName]; ok {
				fmt.Println(sol)
			} else {
				fmt.Println("(no solution available)")
			}
		}

		fmt.Println()
	}
}
