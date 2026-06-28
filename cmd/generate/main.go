package main

import (
	"fmt"
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
		path := filepath.Join(dir, s.ShortName+".go")

		var code string
		pkgName := question.SanitizePkgName(s.ShortName)
		switch s.Kind {
		case "lru-cache":
			code = generateLRUCache(s, pkgName)
		case "codec-tree":
			code = generateCodecTree(s, pkgName)
		case "max-path-tree":
			code = generateMaxPathTree(s, pkgName)
		case "tree":
			code = generateFunctionWithImport(s, pkgName, "github.com/jones/interesting-interview-questions/internal/types")
		case "list":
			code = generateFunctionWithImport(s, pkgName, "github.com/jones/interesting-interview-questions/internal/types")
		default:
			if s.Signature == "" {
				log.Fatalf("question %q: no signature and unknown kind %q", s.ShortName, s.Kind)
			}
			code = generateFunction(s, pkgName)
		}

		if err := os.WriteFile(path, []byte(code), 0644); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("generated %s\n", path)
	}
}

func stubDoc(s question.Spec, name string) string {
	text := strings.ReplaceAll(s.QuestionText, "*/", "* /")
	return "/*\n" + name + " solves " + s.Title + ".\n\n" + strings.TrimSpace(text) + "\n*/"
}

func generateFunction(s question.Spec, pkgName string) string {
	return generateWithPrefix(s, pkgName, "")
}

func generateFunctionWithImport(s question.Spec, pkgName, importPath string) string {
	importBlock := fmt.Sprintf("import \"%s\"\n", importPath)
	return generateWithPrefix(s, pkgName, importBlock)
}

func generateWithPrefix(s question.Spec, pkgName, prefix string) string {
	name, params, returns := parseFuncSig(s.Signature)
	// Replace TreeNode/ListNode with types.TreeNode/types.ListNode for imported types
	params = strings.ReplaceAll(params, "TreeNode", "types.TreeNode")
	params = strings.ReplaceAll(params, "ListNode", "types.ListNode")
	returns = strings.ReplaceAll(returns, "TreeNode", "types.TreeNode")
	returns = strings.ReplaceAll(returns, "ListNode", "types.ListNode")
	// Export the function name (capitalize first letter)
	exportedName := strings.ToUpper(name[:1]) + name[1:]
	var b strings.Builder
	b.WriteString("package ")
	b.WriteString(pkgName)
	b.WriteString("\n\n")
	if prefix != "" {
		b.WriteString(prefix)
		b.WriteString("\n")
	}
	b.WriteString(stubDoc(s, name))
	b.WriteString("\nfunc ")
	b.WriteString(exportedName)
	b.WriteString("(")
	b.WriteString(params)
	b.WriteString(")")
	if returns != "" {
		b.WriteString(" ")
		b.WriteString(returns)
	}
	b.WriteString(" {\n")
	if returns == "" {
		b.WriteString("}\n")
	} else {
		b.WriteString("\treturn ")
		b.WriteString(zeroVals(returns))
		b.WriteString("\n}\n")
	}
	return b.String()
}

func parseFuncSig(sig string) (name, params, returns string) {
	parenIdx := strings.Index(sig, "(")
	if parenIdx < 0 {
		return "", "", ""
	}
	name = strings.TrimSpace(sig[:parenIdx])
	depth := 0
	for i := parenIdx; i < len(sig); i++ {
		switch sig[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				params = strings.TrimSpace(sig[parenIdx+1 : i])
				rest := strings.TrimSpace(sig[i+1:])
				if rest != "" {
					returns = rest
				}
				return
			}
		}
	}
	return "", "", ""
}

func zeroVals(typ string) string {
	if typ == "" {
		return ""
	}
	types := splitReturnTypes(typ)
	vals := make([]string, len(types))
	for i, t := range types {
		vals[i] = zeroValue(t)
	}
	return strings.Join(vals, ", ")
}

func splitReturnTypes(returns string) []string {
	r := strings.TrimSpace(returns)
	if !strings.HasPrefix(r, "(") {
		return []string{r}
	}
	inner := strings.TrimSpace(r[1 : len(r)-1])
	var parts []string
	depth := 0
	start := 0
	for i := 0; i < len(inner); i++ {
		switch inner[i] {
		case '(':
			depth++
		case ')':
			depth--
		case ',':
			if depth == 0 {
				parts = append(parts, strings.TrimSpace(inner[start:i]))
				start = i + 1
			}
		}
	}
	if start < len(inner) {
		parts = append(parts, strings.TrimSpace(inner[start:]))
	}
	return parts
}

func zeroValue(typ string) string {
	typ = strings.TrimSpace(typ)
	if typ == "" {
		return ""
	}
	if strings.HasPrefix(typ, "[]") || strings.HasPrefix(typ, "*") || strings.HasPrefix(typ, "map[") || strings.HasPrefix(typ, "chan ") || strings.HasPrefix(typ, "func(") {
		return "nil"
	}
	switch typ {
	case "string":
		return `""`
	case "bool":
		return "false"
	case "int", "int8", "int16", "int32", "int64":
		return "0"
	case "uint", "uint8", "uint16", "uint32", "uint64":
		return "0"
	case "float32", "float64":
		return "0"
	case "byte", "rune":
		return "0"
	default:
		return "nil"
	}
}

func generateLRUCache(s question.Spec, pkgName string) string {
	return fmt.Sprintf(`package %s

%s
type LRUCache struct{}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{}
}

func (c *LRUCache) Get(key int) int {
	return -1
}

func (c *LRUCache) Put(key int, value int) {}
`, pkgName, stubDoc(s, "LRUCache"))
}

func generateCodecTree(s question.Spec, pkgName string) string {
	return fmt.Sprintf(`package %s

import "github.com/jones/interesting-interview-questions/internal/types"

%s
type Codec struct{}

func Constructor() *Codec {
	return &Codec{}
}

func (c *Codec) Serialize(root *types.TreeNode) string {
	return ""
}

func (c *Codec) Deserialize(data string) *types.TreeNode {
	return nil
}
`, pkgName, stubDoc(s, "Codec"))
}

func generateMaxPathTree(s question.Spec, pkgName string) string {
	return fmt.Sprintf(`package %s

import "github.com/jones/interesting-interview-questions/internal/types"

%s
func MaxPathSum(root *types.TreeNode) int {
	return 0
}
`, pkgName, stubDoc(s, "MaxPathSum"))
}
