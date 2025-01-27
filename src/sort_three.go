package src

import "fmt"

func (s *Stack) Sa() {
	if len(*s) >= 2 {
		(*s)[len(*s)-1], (*s)[len(*s)-2] = (*s)[len(*s)-2], (*s)[len(*s)-1]
	}
}

func (s *Stack) Ra() {
	if len(*s) >= 2 {
		top := (*s)[len(*s)-1]
		*s = append([]int{top}, (*s)[:len(*s)-1]...)
	}
}

func (s *Stack) Rra() {
	if len(*s) >= 2 {
		bottom := (*s)[0]
		*s = append((*s)[1:], bottom)
	}
}

func (s *Stack) SortThree() []string {
	if len(*s) != 3 {
		return nil
	}

	moves := make([]string, 0, 3)
	bottom, middle, top := (*s)[0], (*s)[1], (*s)[2]

	isSorted := func() bool {
		return (*s)[0] < (*s)[1] && (*s)[1] < (*s)[2]
	}

	fmt.Println("Initial stack:", *s)

	// Case 5: `[2,3,1]` -> requires `Ra`
	if top < middle && top < bottom && middle > bottom && middle > top {
		s.Ra()
		moves = append(moves, "ra")
		if isSorted() {
			return moves
		}
	}

	// Case 1: `[2,1,3]` -> requires `Rra` then `Sa`
	if top > middle && middle < bottom && top > bottom {
		s.Rra()
		moves = append(moves, "rra")
		s.Sa()
		moves = append(moves, "sa")
		if isSorted() {
			return moves
		}
	}

	// Case 2: `[3,2,1]` -> requires `Ra` then `Sa`
	if top < middle && middle < bottom {
		s.Ra()
		moves = append(moves, "ra")
		s.Sa()
		moves = append(moves, "sa")
		if isSorted() {
			return moves
		}
	}

	// Case 3: `[3,1,2]` -> requires `Rra`
	if top > middle && middle < bottom && top < bottom {
		s.Rra()
		moves = append(moves, "rra")
		if isSorted() {
			return moves
		}
	}

	// Case 4: `[1,3,2]` -> requires `Sa`
	if top < middle && middle > bottom && bottom < middle {
		s.Sa()
		moves = append(moves, "sa")
		if isSorted() {
			return moves
		}
	}

	fmt.Println("Moves applied:", moves)
	fmt.Println("Final stack:", *s)

	if isSorted() {
		return moves
	}

	// Return nil if the stack is not sorted (as a fallback)
	return nil
}
