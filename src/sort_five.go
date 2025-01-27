package src

func (s *Stack) SortFive(StackB *Stack) []string {
	if len(*s) > 5 {
		return nil
	}

	// StackB := &Stack{}
	var moves []string

	// Move two elements from Stack A to Stack B
	for len(*s) > 3 {
		val, _ := s.Pop()
		StackB.Push(val)
		moves = append(moves, "pb") // Record the move to Stack B
	}

	// Sort the remaining three elements in Stack A
	threeMoves := s.SortThree()
	moves = append(moves, threeMoves...)

	// Handle reinserting elements from Stack B back into Stack A
	for len(*StackB) > 0 {
		topB, _ := StackB.Peek()
		topA, _ := s.Peek()
		bottomA := (*s)[0]        // First element in Stack A
		index1 := (*s)[1]         // Second element in Stack A
		index4 := (*s)[len(*s)-2] // Second-to-last element in Stack A

		// Case: Top of B is the largest and goes on top of Stack A
		if topB > topA {
			val, _ := StackB.Pop()
			s.Push(val)
			moves = append(moves, "pa")
			continue
		}

		// Case: Top of B is the second-largest and should go right below top of Stack A
		if topB < topA && topB > index4 {
			val, _ := StackB.Pop()
			s.Push(val)
			moves = append(moves, "pa")
			s.Sa() // Swap to place it correctly
			moves = append(moves, "sa")
			continue
		}

		// Case: Top of B is smallest, should go to the bottom
		if topB < bottomA {
			val, _ := StackB.Pop()
			s.Push(val)
			moves = append(moves, "pa")
			s.Ra() // Rotate to place it at the bottom
			moves = append(moves, "ra")
			continue
		}

		// Case: Top of B is just above the smallest and should be placed just above bottom
		if topB > bottomA && topB < index1 {
			s.Rra() // Rotate down to position it correctly
			moves = append(moves, "rra")
			val, _ := StackB.Pop()
			s.Push(val)
			moves = append(moves, "pa")
			s.Ra() // Rotate up to restore order
			moves = append(moves, "ra")
			s.Ra()
			moves = append(moves, "ra")
			continue
		}

		// Case: Top of B goes into the middle of Stack A
		if topB > index1 && topB < index4 {
			s.Rra()
			moves = append(moves, "rra")
			s.Rra()
			moves = append(moves, "rra")
			val, _ := StackB.Pop()
			s.Push(val)
			moves = append(moves, "pa")
			// Rotate up to restore sorted order
			s.Ra()
			moves = append(moves, "ra")
			s.Ra()
			moves = append(moves, "ra")
			s.Ra()
			moves = append(moves, "ra")
			continue
		}
	}

	return moves
}
