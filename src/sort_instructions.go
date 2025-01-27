package src

import (
	"fmt"
)

// pa: Push the top element of stackB to stackA
func Pa(stackA, stackB *Stack) error {
	fmt.Println("pa: Push from stackB to stackA")
	value, err := stackB.Pop()
	if err != nil {
		return err
	}
	return stackA.Push(value)
}

// pb: Push the top element of stackA to stackB
func pb(stackA, stackB *Stack) error {
	fmt.Println("pb: Push from stackA to stackB")
	value, err := stackA.Pop()
	if err != nil {
		return err
	}
	return stackB.Push(value)
}

// sa: Swap the top two elements of stackA
func Sa(stackA *Stack) error {
	fmt.Println("sa: Swap top two elements of stackA")
	return stackA.Swap()
}

// sb: Swap the top two elements of stackB
func Sb(stackB *Stack) error {
	fmt.Println("sb: Swap top two elements of stackB")
	return stackB.Swap()
}

// ss: Execute sa and sb simultaneously
func ss(stackA, stackB *Stack) error {
	fmt.Println("ss: Swap top elements of both stackA and stackB")
	if err := Sa(stackA); err != nil {
		return err
	}
	return Sb(stackB)
}

// ra: Rotate stackA (shift up all elements of stackA by 1)
// ra: Rotate stackA (shift up all elements of stackA by 1)
func Ra(stackA *Stack) error {
	fmt.Println("ra: Rotate stackA")
	stackA.RotateTop() // This correctly moves the top to the bottom
	return nil
}

// rb: Rotate stackB (shift up all elements of stackB by 1)
func rb(stackB *Stack) error {
	fmt.Println("rb: Rotate stackB")
	stackB.RotateTop() // This correctly moves the top to the bottom
	return nil
}

// rr: Execute ra and rb simultaneously
func rr(stackA, stackB *Stack) error {
	fmt.Println("rr: Rotate both stackA and stackB")
	Ra(stackA)
	rb(stackB)
	return nil
}

// rra: Reverse rotate stackA (shift down all elements of stackA by 1)
func rra(stackA *Stack) error {
	fmt.Println("rra: Reverse rotate stackA")
	stackA.RotateBottom() // This correctly moves the bottom to the top
	return nil
}

// rrb: Reverse rotate stackB (shift down all elements of stackB by 1)
func rrb(stackB *Stack) error {
	fmt.Println("rrb: Reverse rotate stackB")
	stackB.RotateBottom() // This correctly moves the bottom to the top
	return nil
}

// rrr: Execute rra and rrb simultaneously
func rrr(stackA, stackB *Stack) error {
	fmt.Println("rrr: Reverse rotate both stackA and stackB")
	rra(stackA)
	rrb(stackB)
	return nil
}
