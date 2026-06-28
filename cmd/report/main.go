package main

import (
	"flag"
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
	questionsFilter := flag.String("q", "", "comma-separated short-names to show; 'all' for every question; also shows question text")
	flag.StringVar(questionsFilter, "questions", "", "comma-separated short-names to show; 'all' for every question; also shows question text")
	showSolution := flag.Bool("s", false, "include reference solution")
	flag.BoolVar(showSolution, "solutions", false, "include reference solution")
	easy := flag.Bool("easy", false, "filter by easy difficulty")
	medium := flag.Bool("medium", false, "filter by medium difficulty")
	hard := flag.Bool("hard", false, "filter by hard difficulty")
	expert := flag.Bool("expert", false, "filter by expert difficulty")
	flag.Parse()

	questionsRaw, err := os.ReadFile("questions.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var questions []Question
	if err := yaml.Unmarshal(questionsRaw, &questions); err != nil {
		log.Fatal(err)
	}

	// Build difficulty set from bool flags.
	wantDifficulty := map[string]bool{}
	for _, d := range []struct {
		on   bool
		name string
	}{
		{*easy, "easy"},
		{*medium, "medium"},
		{*hard, "hard"},
		{*expert, "expert"},
	} {
		if d.on {
			wantDifficulty[d.name] = true
		}
	}

	// Parse questions filter.
	var filterNames map[string]bool
	showText := false
	if *questionsFilter != "" {
		showText = true
		filterNames = map[string]bool{}
		for _, s := range strings.Split(*questionsFilter, ",") {
			s = strings.TrimSpace(s)
			if s != "" {
				filterNames[s] = true
			}
		}
	}

	var solutions map[string]string
	if *showSolution {
		solutionsRaw, err := os.ReadFile("solutions.yaml")
		if err != nil {
			log.Fatal(err)
		}
		if err := yaml.Unmarshal(solutionsRaw, &solutions); err != nil {
			log.Fatal(err)
		}
	}

	for _, q := range questions {
		// Filter by difficulty.
		if len(wantDifficulty) > 0 && !wantDifficulty[q.Difficulty] {
			continue
		}
		// Filter by short-name.
		if filterNames != nil {
			if !filterNames[q.ShortName] && !filterNames["all"] && !filterNames["*"] {
				continue
			}
		}

		fmt.Printf("%s (%s, %s) — %s\n", q.ShortName, q.Difficulty, q.Category, q.Title)

		if showText {
			fmt.Println()
			fmt.Println(q.QuestionText)
		}

		if *showSolution {
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
