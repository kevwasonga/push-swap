package src

import (
	"fmt"
	"sort"
	"strconv"
)

// Simplify indices to unique, ordered values
func simplifyIndices(input []int) []int {
	copyInput := make([]int, len(input))
	copy(copyInput, input)
	sort.Ints(copyInput) // sort the inputs in ascending order

	indexMap := make(map[int]int)
	for i, val := range copyInput {
		indexMap[val] = i
	}

	for i, val := range input {
		input[i] = indexMap[val]
	}
	return input
}

// Convert indices to binary strings with uniform length for bitwise sorting
func convertToBinary(indices []int) []string {
	var binaryRepresentations []string
	maxBits := 0

	// Convert to binary and find max bit length
	for _, index := range indices {
		binaryRep := strconv.FormatInt(int64(index), 2)
		binaryRepresentations = append(binaryRepresentations, binaryRep)
		if len(binaryRep) > maxBits {
			maxBits = len(binaryRep)
		}
	}

	// Pad all binary strings to uniform length
	for i, bin := range binaryRepresentations {
		binaryRepresentations[i] = fmt.Sprintf("%0*s", maxBits, bin)

		// fmt.Println("binaryRepresentations!!!!!!!!!!", binaryRepresentations)
	}
	return binaryRepresentations
}
