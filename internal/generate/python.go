package generate

import (
	"strings"

	"github.com/jones/interesting-interview-questions/internal/question"
)

// renderPythonStub returns the self-contained Python stub for a question.
func renderPythonStub(s question.Spec) string {
	var b strings.Builder
	switch s.Kind {
	case "lru-cache":
		b.WriteString(pythonLRUCacheStub(s))
	case "time-based-kv-store":
		b.WriteString(pythonTimeBasedKVStoreStub(s))
	case "codec-tree":
		b.WriteString(pythonCodecTreeStub(s))
	case "max-path-tree":
		b.WriteString(pythonMaxPathTreeStub(s))
	case "tree", "list":
		b.WriteString(pythonTypeDefs(s))
		if s.PythonSignature != "" {
			b.WriteString("\n\n")
			b.WriteString(pythonFunctionStub(s))
		}
	default:
		if s.PythonSignature != "" {
			b.WriteString(pythonFunctionStub(s))
		}
	}
	return b.String()
}

func pythonDoc(s question.Spec) string {
	return "\"\"\"Solve " + s.Title + ". See README.md for the problem statement.\"\"\""
}

func pythonFunctionStub(s question.Spec) string {
	return s.PythonSignature + "\n    " + pythonDoc(s) + "\n    return " + pythonZeroValue(pythonReturnType(s.PythonSignature)) + "\n"
}

func pythonTypeDefs(s question.Spec) string {
	var parts []string
	if needsTreeNode(s) {
		parts = append(parts, "class TreeNode:\n    def __init__(self, val: int = 0, left: 'TreeNode | None' = None, right: 'TreeNode | None' = None):\n        self.val = val\n        self.left = left\n        self.right = right")
	}
	if needsListNode(s) {
		parts = append(parts, "class ListNode:\n    def __init__(self, val: int = 0, next: 'ListNode | None' = None):\n        self.val = val\n        self.next = next")
	}
	return strings.Join(parts, "\n\n")
}

func pythonReturnType(sig string) string {
	idx := strings.Index(sig, "->")
	if idx < 0 {
		return ""
	}
	end := strings.Index(sig[idx:], ":")
	if end < 0 {
		return strings.TrimSpace(sig[idx+2:])
	}
	return strings.TrimSpace(sig[idx+2 : idx+end])
}

func pythonZeroValue(typ string) string {
	typ = strings.TrimSpace(typ)
	if typ == "" {
		return "None"
	}
	if strings.Contains(typ, "|") {
		return "None"
	}
	if strings.HasPrefix(typ, "list") || strings.HasPrefix(typ, "List") {
		return "[]"
	}
	switch typ {
	case "int":
		return "0"
	case "bool":
		return "False"
	case "str":
		return "\"\""
	case "float":
		return "0.0"
	default:
		return "None"
	}
}

func pythonLRUCacheStub(s question.Spec) string {
	return "class LRUCache:\n    " + pythonDoc(s) + "\n    def __init__(self, capacity: int):\n        pass\n\n    def get(self, key: int) -> int:\n        return -1\n\n    def put(self, key: int, value: int) -> None:\n        pass\n"
}

func pythonTimeBasedKVStoreStub(s question.Spec) string {
	return "class TimeMap:\n    " + pythonDoc(s) + "\n    def __init__(self):\n        pass\n\n    def set(self, key: str, value: str, timestamp: int) -> None:\n        pass\n\n    def get(self, key: str, timestamp: int) -> str:\n        return \"\"\n"
}

func pythonCodecTreeStub(s question.Spec) string {
	return "class TreeNode:\n    def __init__(self, val: int = 0, left: 'TreeNode | None' = None, right: 'TreeNode | None' = None):\n        self.val = val\n        self.left = left\n        self.right = right\n\n" +
		"class Codec:\n    " + pythonDoc(s) + "\n    def __init__(self):\n        pass\n\n    def serialize(self, root: TreeNode | None) -> str:\n        return \"\"\n\n    def deserialize(self, data: str) -> TreeNode | None:\n        return None\n"
}

func pythonMaxPathTreeStub(s question.Spec) string {
	return "class TreeNode:\n    def __init__(self, val: int = 0, left: 'TreeNode | None' = None, right: 'TreeNode | None' = None):\n        self.val = val\n        self.left = left\n        self.right = right\n\n" +
		pythonFunctionStub(s)
}
