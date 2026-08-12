package generate

import (
	"strings"

	"github.com/jones/interesting-interview-questions/internal/question"
)

func renderGoMod(s question.Spec, goVersion string) string {
	return "module " + ModuleName(s) + "\n\ngo " + goVersion + "\n"
}

// renderGoStub returns the self-contained Go stub for a question.
func renderGoStub(s question.Spec) string {
	var b strings.Builder
	b.WriteString("package main\n\n")
	switch s.Kind {
	case "lru-cache":
		b.WriteString(lruCacheStub(s))
	case "time-based-kv-store":
		b.WriteString(timeBasedKVStoreStub(s))
	case "codec-tree":
		b.WriteString(codecTreeStub(s))
	case "max-path-tree":
		b.WriteString(maxPathTreeStub(s))
	case "tree", "list":
		b.WriteString(functionStub(s, true))
	default:
		b.WriteString(functionStub(s, false))
	}
	b.WriteString("\nfunc main() {}\n")
	return b.String()
}

func stubDoc(s question.Spec, name string) string {
	return "// " + name + " solves " + s.Title + ". See README.md for the problem statement."
}

func functionStub(s question.Spec, needsType bool) string {
	name, params, returns := parseFuncSig(s.Signature)
	exported := strings.ToUpper(name[:1]) + name[1:]
	var b strings.Builder
	if needsType {
		b.WriteString(typeDefs(s))
		b.WriteString("\n\n")
	}
	b.WriteString(stubDoc(s, exported))
	b.WriteString("\nfunc ")
	b.WriteString(exported)
	b.WriteString("(")
	b.WriteString(params)
	b.WriteString(")")
	if returns != "" {
		b.WriteString(" ")
		b.WriteString(returns)
	}
	b.WriteString(" {\n")
	if returns != "" {
		b.WriteString("\treturn ")
		b.WriteString(zeroVals(returns))
		b.WriteString("\n")
	}
	b.WriteString("}\n")
	return b.String()
}

func typeDefs(s question.Spec) string {
	var parts []string
	if needsTreeNode(s) {
		parts = append(parts, "type TreeNode struct {\n\tVal   int\n\tLeft  *TreeNode\n\tRight *TreeNode\n}")
	}
	if needsListNode(s) {
		parts = append(parts, "type ListNode struct {\n\tVal  int\n\tNext *ListNode\n}")
	}
	return strings.Join(parts, "\n\n")
}

func needsTreeNode(s question.Spec) bool {
	return s.Kind == "tree" || s.Kind == "codec-tree" || s.Kind == "max-path-tree" ||
		strings.Contains(s.Signature, "TreeNode")
}

func needsListNode(s question.Spec) bool {
	return s.Kind == "list" || strings.Contains(s.Signature, "ListNode")
}

func lruCacheStub(s question.Spec) string {
	return stubDoc(s, "LRUCache") + `
type LRUCache struct{}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{}
}

func (c *LRUCache) Get(key int) int {
	return -1
}

func (c *LRUCache) Put(key int, value int) {}
`
}

func timeBasedKVStoreStub(s question.Spec) string {
	return stubDoc(s, "TimeMap") + `
type TimeMap struct{}

func Constructor() *TimeMap {
	return &TimeMap{}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {}

func (t *TimeMap) Get(key string, timestamp int) string {
	return ""
}
`
}

func codecTreeStub(s question.Spec) string {
	return "type TreeNode struct {\n\tVal   int\n\tLeft  *TreeNode\n\tRight *TreeNode\n}\n\n" +
		stubDoc(s, "Codec") + `
type Codec struct{}

func Constructor() *Codec {
	return &Codec{}
}

func (c *Codec) Serialize(root *TreeNode) string {
	return ""
}

func (c *Codec) Deserialize(data string) *TreeNode {
	return nil
}
`
}

func maxPathTreeStub(s question.Spec) string {
	return "type TreeNode struct {\n\tVal   int\n\tLeft  *TreeNode\n\tRight *TreeNode\n}\n\n" +
		stubDoc(s, "MaxPathSum") + `
func MaxPathSum(root *TreeNode) int {
	return 0
}
`
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
