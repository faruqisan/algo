package blpts

import (
	"github.com/golang-collections/collections/stack"
)

var (
	pair = map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}
)

// IsBalanced function check wheter given sets of parentheses are balanced
// check the unit test for example of parameters
func IsBalanced(parentheses []string) bool {
	var (
		result      bool
		startingPts = stack.New()
	)

	for _, v := range parentheses {
		// check if current value of parentheses is starting "{", "(", "["
		if isStartingParentheses(v) {
			// push to stack
			startingPts.Push(v)
		} else {
			// check if there's already pushed starting pts
			if startingPts.Len() < 1 {
				return false
			}
			// get nearest pts
			lastParentheses := startingPts.Pop()
			if !isMatchingPair(lastParentheses.(string), v) {
				return false
			}
		}
	}

	// check if there's no starting parentheses left
	if startingPts.Len() < 1 {
		result = true
	}

	return result
}

func isStartingParentheses(s string) bool {
	var (
		result bool
	)

	if s == "{" || s == "(" || s == "[" {
		result = true
	}

	return result

}

func isMatchingPair(s1, s2 string) bool {
	var (
		result bool
	)

	if val, ok := pair[s1]; ok {
		if val == s2 {
			return true
		}
	}

	return result
}
