// LeetCode 234 - Palindrome Linked List

package stacks

import "fmt"

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
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

func maxDepth(s string) int {
	stack := Stack{}
	countDepth := 0
	tmpCloseDepth := 0

	for _, char := range s {
		if char == '(' {
			top := stack.Top()
			if top == "(" || top == "" {
				tmpCloseDepth++
			}
			stack.Push(string(char))
		} else if char == ')' {
			top := stack.Top()
			if top == "(" {
				if countDepth < tmpCloseDepth {
					countDepth = tmpCloseDepth
				}
			}

			tmpCloseDepth--
			stack.Pop()
		}
	}

	return countDepth
}
