package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jones/interesting-interview-questions/internal/question"
)

func main() {
	specs, err := question.LoadAll("questions.yaml")
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range specs {
		dir := filepath.Join("questions", s.Difficulty, s.ShortName)
		testPath := filepath.Join(dir, s.ShortName+"_test.go")

		content, err := os.ReadFile(testPath)
		if err != nil {
			fmt.Printf("SKIP %s: %v\n", testPath, err)
			continue
		}

		newContent := transformTestFile(string(content), s)
		formatted, err := format.Source([]byte(newContent))
		if err != nil {
			fmt.Printf("FORMAT ERROR %s: %v\n", testPath, err)
			continue
		}

		if !bytes.Equal(content, formatted) {
			if err := os.WriteFile(testPath, formatted, 0644); err != nil {
				fmt.Printf("WRITE ERROR %s: %v\n", testPath, err)
				continue
			}
			fmt.Printf("UPDATED %s\n", testPath)
		} else {
			fmt.Printf("UNCHANGED %s\n", testPath)
		}
	}
}

func transformTestFile(content string, spec question.Spec) string {
	pkgName := question.SanitizePkgName(spec.ShortName)

	// Replace package declaration
	content = strings.Replace(content, "package main", "package "+pkgName+"_test", 1)

	// Fix existing double-prefixes from prior runs before doing new replacements
	content = strings.ReplaceAll(content, "types.types.TreeNode", "types.TreeNode")
	content = strings.ReplaceAll(content, "types.types.ListNode", "types.ListNode")
	content = strings.ReplaceAll(content, pkgName+"."+pkgName+".", pkgName+".")
	// For method calls on a receiver: varName.pkgName.Method( -> varName.Method(
	for _, fn := range getSpecialFuncNames(spec.Kind) {
		content = strings.ReplaceAll(content, "."+pkgName+"."+fn+"(", "."+fn+"(")
	}

	// Determine imports needed
	needsTypes := false
	needsTestUtil := false
	needsSort := false
	needsReflect := false

	if strings.Contains(content, "TreeNode") {
		needsTypes = true
	}
	if strings.Contains(content, "ListNode") {
		needsTypes = true
	}
	if strings.Contains(content, "sort.Ints") || strings.Contains(content, "sort.Strings") || strings.Contains(content, "sort.Slice") {
		needsSort = true
	}
	if strings.Contains(content, "reflect.DeepEqual") {
		needsReflect = true
	}
	hasTestutil := strings.Contains(content, "testutil.")
	hasSortPattern := strings.Contains(content, "sort.Ints(got)") || strings.Contains(content, "sort.Strings(got)")
	hasReflectPattern := strings.Contains(content, "if !reflect.DeepEqual(got, want) {") || strings.Contains(content, "if !reflect.DeepEqual(tt.s, tt.want) {") || strings.Contains(content, "if !reflect.DeepEqual(got, tt.want) {")
	if hasTestutil || (hasSortPattern && hasReflectPattern) {
		needsTestUtil = true
	}

	importPath := fmt.Sprintf("\"github.com/jones/interesting-interview-questions/questions/%s/%s\"", spec.Difficulty, spec.ShortName)

	var imports []string
	imports = append(imports, "\"testing\"")
	imports = append(imports, importPath)
	if needsTypes {
		imports = append(imports, "\"github.com/jones/interesting-interview-questions/internal/types\"")
	}
	if needsSort {
		imports = append(imports, "\"sort\"")
	}
	if needsReflect {
		imports = append(imports, "\"reflect\"")
	}
	if needsTestUtil {
		imports = append(imports, "\"github.com/jones/interesting-interview-questions/internal/testutil\"")
	}

	importBlock := "import (\n\t" + strings.Join(imports, "\n\t") + "\n)"
	content = replaceImportBlock(content, importBlock)

	// Replace type references (idempotent: only prefix standalone references)
	if needsTypes {
		content = fixTypeReferences(content)
	}

	// Replace function calls with package prefix
	specialFuncs := getSpecialFuncNames(spec.Kind)
	if len(specialFuncs) > 0 {
		for _, fn := range specialFuncs {
			content = replaceFuncCalls(content, pkgName, fn, fn)
			// Fix unexported method calls: codec.serialize( -> codec.Serialize(
			if len(fn) > 0 {
				unexported := strings.ToLower(fn[:1]) + fn[1:]
				if unexported != fn {
					content = fixMethodCase(content, unexported, fn)
				}
			}
		}
	} else {
		origFuncName := extractOrigFuncName(spec.Signature)
		exportedFuncName := extractFuncName(spec.Signature)
		if origFuncName != "" {
			content = replaceFuncCalls(content, pkgName, origFuncName, exportedFuncName)
		}
	}

	// Replace sort/reflect with testutil helpers
	if needsTestUtil {
		content = strings.ReplaceAll(content, "sort.Ints(got)", "testutil.AssertEqualSliceInt(t, got, want)")
		content = strings.ReplaceAll(content, "sort.Strings(got)", "testutil.AssertEqualSliceString(t, got, want)")
		content = strings.ReplaceAll(content, "sort.Ints(want)", "")
		content = strings.ReplaceAll(content, "sort.Strings(want)", "")

		content = strings.ReplaceAll(content, "if !reflect.DeepEqual(got, want) {", "testutil.AssertEqual(t, got, want)")
		content = strings.ReplaceAll(content, "if !reflect.DeepEqual(tt.s, tt.want) {", "testutil.AssertEqual(t, tt.s, tt.want)")
		content = strings.ReplaceAll(content, "if !reflect.DeepEqual(got, tt.want) {", "testutil.AssertEqual(t, got, tt.want)")
		content = strings.ReplaceAll(content, "if got != tt.want {", "testutil.AssertEqual(t, got, tt.want)")
	}

	// Clean up empty lines
	content = strings.ReplaceAll(content, "\n\n\n", "\n\n")

	return content
}

func fixTypeReferences(content string) string {
	// Protect already-correct references, then add prefix only to standalone ones.
	// Use null byte markers which are safe in Go source code.
	content = strings.ReplaceAll(content, "types.TreeNode", "\x00TYPES_TREE\x00")
	content = strings.ReplaceAll(content, "types.ListNode", "\x00TYPES_LIST\x00")

	content = strings.ReplaceAll(content, "*TreeNode", "*types.TreeNode")
	content = strings.ReplaceAll(content, "TreeNode", "types.TreeNode")
	content = strings.ReplaceAll(content, "*ListNode", "*types.ListNode")
	content = strings.ReplaceAll(content, "ListNode", "types.ListNode")

	content = strings.ReplaceAll(content, "\x00TYPES_TREE\x00", "types.TreeNode")
	content = strings.ReplaceAll(content, "\x00TYPES_LIST\x00", "types.ListNode")
	return content
}

func replaceImportBlock(content, newImportBlock string) string {
	lines := strings.Split(content, "\n")
	var result []string
	inImport := false
	importDone := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "import (" && !importDone {
			inImport = true
			result = append(result, newImportBlock)
			continue
		}
		if inImport {
			if trimmed == ")" {
				inImport = false
				importDone = true
			}
			continue
		}
		if strings.HasPrefix(trimmed, "import \"") && !importDone {
			result = append(result, newImportBlock)
			importDone = true
			continue
		}
		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

func extractFuncName(sig string) string {
	parenIdx := strings.Index(sig, "(")
	if parenIdx < 0 {
		return ""
	}
	name := strings.TrimSpace(sig[:parenIdx])
	return strings.ToUpper(name[:1]) + name[1:]
}

func extractOrigFuncName(sig string) string {
	parenIdx := strings.Index(sig, "(")
	if parenIdx < 0 {
		return ""
	}
	return strings.TrimSpace(sig[:parenIdx])
}

func getSpecialFuncNames(kind string) []string {
	switch kind {
	case "lru-cache":
		return []string{"Constructor", "Get", "Put"}
	case "codec-tree":
		return []string{"Constructor", "Serialize", "Deserialize"}
	case "max-path-tree":
		return []string{"MaxPathSum"}
	default:
		return nil
	}
}

func fixMethodCase(content, unexported, exported string) string {
	// Replace .unexported( with .Exported( for method calls on a receiver
	// e.g., codec.serialize(root) -> codec.Serialize(root)
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "//") {
			continue
		}
		lines[i] = strings.ReplaceAll(line, "."+unexported+"(", "."+exported+"(")
	}
	return strings.Join(lines, "\n")
}

func replaceFuncCalls(content, pkgName, origFuncName, exportedFuncName string) string {
	search := origFuncName + "("
	replace := pkgName + "." + exportedFuncName + "("

	lines := strings.Split(content, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "func ") {
			continue
		}

		// Skip lines that already have the package prefix
		if strings.Contains(line, pkgName+"."+origFuncName+"(") ||
			strings.Contains(line, pkgName+"."+exportedFuncName+"(") {
			continue
		}

		// Only replace standalone occurrences: not preceded by a word char, dot, or underscore
		var result strings.Builder
		last := 0
		for {
			idx := strings.Index(line[last:], search)
			if idx < 0 {
				break
			}
			pos := last + idx
			if pos > 0 {
				prev := line[pos-1]
				if (prev >= 'a' && prev <= 'z') || (prev >= 'A' && prev <= 'Z') ||
					(prev >= '0' && prev <= '9') || prev == '.' || prev == '_' {
					// Already prefixed or method call — copy verbatim
					result.WriteString(line[last : pos+len(search)])
					last = pos + len(search)
					continue
				}
			}
			result.WriteString(line[last:pos])
			result.WriteString(replace)
			last = pos + len(search)
		}
		result.WriteString(line[last:])
		lines[i] = result.String()
	}
	return strings.Join(lines, "\n")
}
