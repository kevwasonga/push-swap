package src

import "errors"

// Stack represents a stack of integers.
type Stack []int

// default error message to return in case of an error
const errMessage = "Error"

// Peek retrieves the top element of the stack without removing it.
// Returns an error if the stack is empty.
func (s *Stack) Peek() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("stack is empty")
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
func (s *Stack) Push(value int) error {
	maxTop := len(*s) + 1 // Define maxTop based on the current length of the stack

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
