package src

func (s *Stack) RotateTop() {
	if s.IsEmpty() {
		return
	}

	// Step 1: Hold the top element
	top, _ := s.Pop()

	// Step 2: Shift remaining elements to a temp stack
	tempStack := Stack{}
	for !s.IsEmpty() {
		value, _ := s.Pop()
		tempStack.Push(value)
	}

	// Step 3: Place held top element at the bottom
	tempStack.Push(top)

	// Step 4: Restore elements back to the original stack
	for !tempStack.IsEmpty() {
		value, _ := tempStack.Pop()
		s.Push(value)
	}
}

func (s *Stack) RotateBottom() {
	if s.IsEmpty() {
		return
	}

	// Step 1: Move all elements except the bottom one to a temp stack
	tempStack := Stack{}
	for len(*s) > 1 {
		value, _ := s.Pop()
		tempStack.Push(value)
	}

	// Step 2: The bottom element is the last element in `s`
	bottom := (*s)[0]
	*s = (*s)[:0] // Clear the original stack

	// Step 3: Move the elements back from tempStack to original stack first
	for !tempStack.IsEmpty() {
		value, _ := tempStack.Pop()
		s.Push(value)
	}

	// Step 4: Finally, push the former bottom element onto the top of the stack
	s.Push(bottom)
}
