package src

import "fmt"

func (stackA *Stack) RadixSort(stackB *Stack) []string {
	var moves []string
	maxBits := getMaxBits(*stackA) // Get the number of bits required for the largest number

	for bitPos := 0; bitPos < maxBits; bitPos++ {
		numElements := len(*stackA)
		lastZeroIndex := findLastZero(*stackA, bitPos)

		for i := 0; i < numElements; i++ {
			top := (*stackA)[0]
			*stackA = (*stackA)[1:] // Pop from stackA

			if (top>>bitPos)&1 == 0 {
				*stackB = append([]int{top}, *stackB...) // Push to top of stackB
				moves = append(moves, "pb")
			} else {
				if i < lastZeroIndex {
					*stackA = append(*stackA, top) // Rotate within stackA
					moves = append(moves, "ra")
				} else {
					*stackA = append(*stackA, top) // Append without rotation
				}
			}
		}

		// Optimize "pa" operations
		numElementsB := len(*stackB)
		lastOneIndex := findLastOne(*stackB, bitPos)

		for i := 0; i < numElementsB; i++ {
			top := (*stackB)[0]
			*stackB = (*stackB)[1:] // Pop from stackB

			if (top>>bitPos)&1 == 1 {
				*stackA = append([]int{top}, *stackA...) // Push to top of stackA
				moves = append(moves, "pa")
			} else if i < lastOneIndex {
				*stackB = append(*stackB, top) // Rotate within stackB
				moves = append(moves, "rb")
			} else {
				*stackB = append(*stackB, top) // Append without rotation
			}
		}
	}

	// Final cleanup
	for len(*stackB) > 0 {
		top := (*stackB)[0]
		*stackB = (*stackB)[1:]                  // Pop from stackB
		*stackA = append([]int{top}, *stackA...) // Push to top of stackA
		moves = append(moves, "pa")
	}

	fmt.Println("TOTAL MOVES USED TO SORT :", len(moves))
	return moves
}

func findLastZero(stack Stack, bitPos int) int {
	lastZeroIndex := -1
	for i, num := range stack {
		if (num>>bitPos)&1 == 0 {
			lastZeroIndex = i
		}
	}
	return lastZeroIndex
}

func findLastOne(stack Stack, bitPos int) int {
	lastOneIndex := -1
	for i, num := range stack {
		if (num>>bitPos)&1 == 1 {
			lastOneIndex = i
		}
	}
	return lastOneIndex
}

// getMaxBits calculates the maximum number of bits needed for binary representation
// of the largest number in the stack.
func getMaxBits(stack Stack) int {
	maxVal := 0
	for _, num := range stack {
		if num > maxVal {
			maxVal = num
		}
	}
	bitCount := 0
	for maxVal > 0 {
		bitCount++
		maxVal >>= 1
	}
	return bitCount
}
