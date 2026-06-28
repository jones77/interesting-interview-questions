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

type QuestionOutput struct {
	ShortName    string `yaml:"short_name"`
	Title        string `yaml:"title"`
	QuestionText string `yaml:"question_text,omitempty"`
	Hint         string `yaml:"hint,omitempty"`
	Difficulty   string `yaml:"difficulty"`
	Category     string `yaml:"category"`
	Solution     string `yaml:"solution,omitempty"`
}

func main() {
	showQuestion := flag.Bool("q", false, "include question text")
	flag.BoolVar(showQuestion, "questions", false, "include question text")
	showSolution := flag.Bool("s", false, "include reference solution (requires -f)")
	flag.BoolVar(showSolution, "solution", false, "include reference solution (requires -f)")
	showAll := flag.Bool("a", false, "include both question text and solution")
	flag.BoolVar(showAll, "all", false, "include both question text and solution")
	var listMode bool
	flag.BoolVar(&listMode, "l", false, "list questions (default)")
	flag.BoolVar(&listMode, "list", false, "list questions (default)")
	filterFlag := flag.String("f", "", "comma-separated short-names to include (default: all)")
	flag.StringVar(filterFlag, "filter", "", "comma-separated short-names to include (default: all)")
	easy := flag.Bool("easy", false, "filter by easy difficulty")
	flag.BoolVar(easy, "e", false, "filter by easy difficulty")
	medium := flag.Bool("medium", false, "filter by medium difficulty")
	flag.BoolVar(medium, "m", false, "filter by medium difficulty")
	hard := flag.Bool("hard", false, "filter by hard difficulty")
	flag.BoolVar(hard, "d", false, "filter by hard difficulty")
	expert := flag.Bool("expert", false, "filter by expert difficulty")
	flag.BoolVar(expert, "x", false, "filter by expert difficulty")
	flag.Parse()

	hasFilter := *filterFlag != ""

	if *showAll && hasFilter {
		log.Fatal("-a and -f are mutually exclusive")
	}
	if *showAll && *showSolution {
		log.Fatal("-a and -s are mutually exclusive")
	}
	if *showSolution && !hasFilter {
		log.Fatal("-s requires -f")
	}

	if *showAll {
		*showQuestion = true
		*showSolution = true
	}

	// With -f and no -q or -s, show question text only.
	if hasFilter && !*showQuestion && !*showSolution {
		*showQuestion = true
	}

	questionsRaw, err := os.ReadFile("questions.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var questions []Question
	if err := yaml.Unmarshal(questionsRaw, &questions); err != nil {
		log.Fatal(err)
	}

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

	var filterNames map[string]bool
	if hasFilter {
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

	verbose := hasFilter || *showQuestion || *showAll

	var matched []QuestionOutput
	for _, q := range questions {
		if len(wantDifficulty) > 0 && !wantDifficulty[q.Difficulty] {
			continue
		}
		if filterNames != nil && !filterNames[q.ShortName] {
			continue
		}

		out := QuestionOutput{
			ShortName:  q.ShortName,
			Title:      q.Title,
			Difficulty: q.Difficulty,
			Category:   q.Category,
		}
		if *showQuestion {
			out.QuestionText = q.QuestionText
			out.Hint = q.Hint
		}
		if *showSolution {
			out.Solution = solutions[q.ShortName]
		}
		matched = append(matched, out)
	}

	if verbose {
		outRaw, err := yaml.Marshal(matched)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(outRaw))
	} else {
		for _, m := range matched {
			fmt.Printf("%s (%s, %s) — %s\n", m.ShortName, m.Difficulty, m.Category, m.Title)
		}
	}
}
