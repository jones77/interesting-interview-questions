package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jones/interesting-interview-questions/internal/question"
	"gopkg.in/yaml.v3"
)

type QuestionOutput struct {
	ShortName    string `yaml:"short_name"`
	Title        string `yaml:"title"`
	Difficulty   string `yaml:"difficulty"`
	Category     string `yaml:"category"`
	Frequency    string `yaml:"frequency"`
	URL          string `yaml:"url,omitempty"`
	QuestionText string `yaml:"question_text,omitempty"`
	Solution     string `yaml:"solution,omitempty"`
}

func main() {
	var (
		showQuestions bool
		showSolution  bool
		showAll       bool
		filterFlag    string
		diffFlag      string
	)

	type flagDef struct{ short, long, usage string }
	var flagDefs []flagDef

	regBool := func(short, long, usage string, p *bool) {
		flagDefs = append(flagDefs, flagDef{short, long, usage})
		flag.BoolVar(p, short, false, "")
		flag.BoolVar(p, long, false, "")
	}
	regString := func(short, long, usage string, p *string) {
		flagDefs = append(flagDefs, flagDef{short, long, usage})
		flag.StringVar(p, short, "", "")
		flag.StringVar(p, long, "", "")
	}

	allQuestions, err := question.LoadAll("questions.yaml")
	if err != nil {
		log.Fatal(err)
	}

	counts := map[string]int{}
	for _, q := range allQuestions {
		counts[q.Difficulty]++
	}
	parts := []string{fmt.Sprintf("%d total", len(allQuestions))}
	for _, d := range []string{"easy", "medium", "hard", "expert"} {
		if c := counts[d]; c > 0 {
			parts = append(parts, fmt.Sprintf("%d %s", c, d))
		}
	}
	summary := strings.Join(parts, ", ")

	regBool("a", "all", "both question text and solution (YAML) for every question", &showAll)
	regString("d", "difficulty", "comma-separated difficulty filter (easy,medium,hard,expert)", &diffFlag)
	regString("f", "filter", "filter by comma-separated short-names; shows question text (YAML)", &filterFlag)
	regBool("q", "questions", fmt.Sprintf("question text (YAML) for all questions (%s)", summary), &showQuestions)
	regBool("s", "solution", "include solution with -f (YAML); requires -f", &showSolution)
	flagDefs = append(flagDefs, flagDef{"h", "help", "show this help"})

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		for _, f := range flagDefs {
			fmt.Fprintf(flag.CommandLine.Output(), "  -%s, --%s\n    \t%s\n", f.short, f.long, f.usage)
		}
	}

	flag.Parse()

	if showAll && filterFlag != "" {
		log.Fatal("-a and -f are mutually exclusive")
	}
	if showAll && showSolution {
		log.Fatal("-a and -s are mutually exclusive")
	}
	if showSolution && filterFlag == "" {
		log.Fatal("-s requires -f")
	}

	if showAll {
		showQuestions = true
		showSolution = true
	}

	if filterFlag != "" && !showQuestions {
		showQuestions = true
	}

	wantDifficulty := map[string]bool{}
	for _, s := range strings.Split(diffFlag, ",") {
		s = strings.TrimSpace(s)
		if s != "" {
			wantDifficulty[s] = true
		}
	}

	var filterNames map[string]bool
	if filterFlag != "" {
		filterNames = map[string]bool{}
		for _, s := range strings.Split(filterFlag, ",") {
			s = strings.TrimSpace(s)
			if s != "" {
				filterNames[s] = true
			}
		}
	}

	verbose := filterFlag != "" || showQuestions || showAll

	var matched []QuestionOutput
	for _, s := range allQuestions {
		if len(wantDifficulty) > 0 && !wantDifficulty[s.Difficulty] {
			continue
		}
		if filterNames != nil && !filterNames[s.ShortName] {
			continue
		}
		out := QuestionOutput{
			ShortName:  s.ShortName,
			Title:      s.Title,
			Difficulty: s.Difficulty,
			Category:   s.Category,
			Frequency:  s.Frequency,
			URL:        s.URL,
		}
		if showQuestions {
			out.QuestionText = s.QuestionText
		}
		if showSolution {
			out.Solution = s.Solution
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
			fmt.Printf("%s (%s/%s, %s) — %s\n", m.ShortName, m.Difficulty, m.Frequency, m.Category, m.Title)
		}
	}
}
