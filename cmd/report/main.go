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
	showQuestion := flag.Bool("q", false, "include question text")
	flag.BoolVar(showQuestion, "questions", false, "include question text")
	showSolution := flag.Bool("s", false, "include reference solution")
	flag.BoolVar(showSolution, "solutions", false, "include reference solution")
	filterFlag := flag.String("f", "", "comma-separated short-names to include (default: all)")
	flag.StringVar(filterFlag, "filter", "", "comma-separated short-names to include (default: all)")
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

	// Parse name filter.
	var filterNames map[string]bool
	if *filterFlag != "" {
		filterNames = map[string]bool{}
		for _, s := range strings.Split(*filterFlag, ",") {
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
		if len(wantDifficulty) > 0 && !wantDifficulty[q.Difficulty] {
			continue
		}
		if filterNames != nil && !filterNames[q.ShortName] {
			continue
		}

		fmt.Printf("%s (%s, %s) — %s\n", q.ShortName, q.Difficulty, q.Category, q.Title)

		if *showQuestion {
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
