package stacks

import "fmt"

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Top() string {
	if s.IsEmpty() {
		return ""
	}
	index := len(*s) - 1
	fmt.Println((*s)[index])
	return (*s)[index]
}

type Stack []string

func isValid(s string) bool {
	stack := Stack{}

	for _, char := range s {
		switch char {
		case '(':
			stack.Push(string(char))
		case '[':
			stack.Push(string(char))
		case '{':
			stack.Push(string(char))
		case ')':
			top := stack.Top()
			if top == "(" {
				stack.Pop()
			} else {
				return false
			}
		case ']':
			top := stack.Top()
			if top == "[" {
				stack.Pop()
			} else {
				return false
			}
		case '}':
			top := stack.Top()
			if top == "{" {
				stack.Pop()
			} else {
				return false
			}
		}
	}

	if stack.IsEmpty() {
		return true
	}

	return false
}

// 20. Valid Parentheses - leetCode

func main() {
	// s := "()"
	s := "()[]{}"
	// s := "(]"
	isValid(s)

}
