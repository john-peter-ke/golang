package leetcode

import (
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	str := strconv.Itoa(x)
	var operator string
	if str[:1] == "-" {
		operator = str[:1]
		str = str[1:]
	}

	stack := NewStack()
	for _, r := range str {
		stack.Push(r)
	}

	var b strings.Builder
	for !stack.IsEmpty() {
		b.WriteString(string(stack.Pop()))
	}

	num, err := strconv.Atoi(b.String())
	if err != nil {
		return 0
	}

	if num > math.MaxInt32 {
		return 0
	}

	if operator != "" {
		return -num
	}

	return num
}

type stack struct {
	values []rune
}

func NewStack() *stack {
	return &stack{
		values: make([]rune, 0),
	}
}

func (s *stack) size() int {
	return len(s.values)
}

func (s *stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *stack) Push(value rune) {
	s.values = append(s.values, value)
}

func (s *stack) Pop() rune {
	if s.IsEmpty() {
		return 0
	}

	index := s.size() - 1
	var value = s.values[index]
	s.values = s.values[:index]
	return value
}
