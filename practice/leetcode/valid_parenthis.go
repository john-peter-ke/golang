package leetcode

import "github.com/mucunga90/go/datastructure/stack"

func isValidParenthesis(s string) bool {
	stk := stack.NewStack[string]()

	for i := 0; i < len(s); i++ {
		if stk.IsEmpty() {
			stk.Push(swapPara(string(s[i])))
		} else if in(string(s[i])) {
			stk.Push(swapPara(string(s[i])))

		} else if out(string(s[i])) {
			if stk.Pop() != string(s[i]) {
				return false
			}
		}
	}

	if !stk.IsEmpty() {
		return false
	}

	return true
}

func in(c string) bool {
	switch c {
	case "(":
		return true
	case "{":
		return true
	case "[":
		return true
	default:
		return false
	}
}

func out(c string) bool {
	switch c {
	case ")":
		return true
	case "}":
		return true
	case "]":
		return true
	default:
		return false
	}
}

func swapPara(c string) string {
	switch c {
	case "(":
		return ")"
	case "{":
		return "}"
	case "[":
		return "]"
	default:
		return ""
	}
}
