package src

import "errors"

// Swap swaps the first two elements of the stack.
func (s *Stack) Swap() error {
	// Check if there are at least two elements to swap
	if len(*s) < 2 {
		return errors.New(errMessage)
	}

	// Swap the top two elements
	(*s)[len(*s)-1], (*s)[len(*s)-2] = (*s)[len(*s)-2], (*s)[len(*s)-1]
	return nil
}
