package src

import "errors"

// Stack represents a stack of integers.
type Stack []int

// default error message to return in case of an error
const errMessage = "Error"

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New(errMessage)
	}
	return (*s)[len(*s)-1], nil
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// check if stack is full
func (s *Stack) IsFull(maxTop int) bool {
	return len(*s) >= maxTop
}

// Push adds a value to the stack if it's not full.
// It returns an error message if the stack is full.
func (s *Stack) Push(value int, maxTop int) error {
	if s.IsFull(maxTop) {
		return errors.New(errMessage)
	}
	*s = append(*s, value)
	return nil
}

func (s *Stack) Pop() (int, error) {
	// Check if the stack is empty
	if s.IsEmpty() {
		return 0, errors.New(errMessage)
	}

	// Get the top value
	top := (*s)[len(*s)-1]

	// Update the stack to remove the top value
	*s = (*s)[:len(*s)-1] // Slice the stack to remove the top element

	return top, nil // Return the top value and no error
}
