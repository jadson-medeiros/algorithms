// LeetCode 234 - Palindrome Linked List

package stacks

type ListNode struct {
	Val  int
	Next *ListNode
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func (s *Stack) Top() int {
	if s.IsEmpty() {
		return -1
	}
	index := len(*s) - 1

	return (*s)[index]
}

type Stack []int

func isPalindrome(head *ListNode) bool {
	stack := Stack{}
	node := head
	for node != nil {
		stack.Push(node.Val)
		node = node.Next
	}

	node = head
	for node != nil {
		if node.Val == stack.Top() {
			node = node.Next
			stack.Pop()
			continue
		}
		return false
	}
	return true
}
