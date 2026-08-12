package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jones/interesting-interview-questions/internal/question"
)

// RootGoVersion returns the go directive version from the module in rootDir.
func RootGoVersion(rootDir string) (string, error) {
	raw, err := os.ReadFile(filepath.Join(rootDir, "go.mod"))
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(raw), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "go ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "go")), nil
		}
	}
	return "", fmt.Errorf("no go directive in %s", filepath.Join(rootDir, "go.mod"))
}

// QuestionDir returns the canonical directory for a question.
func QuestionDir(rootDir string, s question.Spec) string {
	return filepath.Join(rootDir, "questions", s.Difficulty, s.ShortName)
}

// EmitAll regenerates every generated artifact for all questions in questions.yaml.
func EmitAll(rootDir string) ([]question.Spec, error) {
	specs, err := question.LoadAll(filepath.Join(rootDir, "questions.yaml"))
	if err != nil {
		return nil, err
	}
	goVersion, err := RootGoVersion(rootDir)
	if err != nil {
		return nil, err
	}
	for _, s := range specs {
		if err := EmitSpec(rootDir, goVersion, s); err != nil {
			return nil, err
		}
	}
	return specs, nil
}

// EmitSpec writes the README.md, go.mod, Go stub, and Python stub for one question.
func EmitSpec(rootDir, goVersion string, s question.Spec) error {
	dir := QuestionDir(rootDir, s)
	goDir := filepath.Join(dir, "go")
	pyDir := filepath.Join(dir, "python")
	if err := os.MkdirAll(goDir, 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(pyDir, 0o755); err != nil {
		return err
	}
	files := map[string]string{
		filepath.Join(dir, "README.md"):           renderREADME(s),
		filepath.Join(goDir, "go.mod"):            renderGoMod(s, goVersion),
		filepath.Join(goDir, s.ShortName+".go"):   renderGoStub(s),
		filepath.Join(pyDir, pythonModuleName(s)): renderPythonStub(s),
	}
	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			return err
		}
	}
	return nil
}

// pythonModuleName returns the Python module filename for a question.
func pythonModuleName(s question.Spec) string {
	return strings.ReplaceAll(s.ShortName, "-", "_") + ".py"
}

// ModuleName returns the bare module name for a question's Go module.
func ModuleName(s question.Spec) string {
	return question.SanitizePkgName(s.ShortName)
}
