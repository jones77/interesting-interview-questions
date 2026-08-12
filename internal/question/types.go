package question

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func SanitizePkgName(name string) string {
	return strings.ReplaceAll(name, "-", "")
}

type Spec struct {
	ShortName       string `yaml:"short_name"`
	Title           string `yaml:"title"`
	Difficulty      string `yaml:"difficulty"`
	Category        string `yaml:"category"`
	Frequency       string `yaml:"frequency"`
	Kind            string `yaml:"kind,omitempty"`
	Signature       string `yaml:"signature,omitempty"`
	PythonSignature string `yaml:"python_signature,omitempty"`
	QuestionText    string `yaml:"question_text"`
	URL             string `yaml:"url,omitempty"`
	Solution        string `yaml:"solution,omitempty"`
}

func LoadAll(filename string) ([]Spec, error) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var specs []Spec
	if err := yaml.Unmarshal(raw, &specs); err != nil {
		return nil, err
	}
	return specs, nil
}
