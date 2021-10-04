/*1. Create a simple stack which can hold a fixed number of ints.
It does not have to grow beyond this limit. Define push – put something on the stack
 – and pop – retrieve something from the stack – functions.
The stack should be a LIFO (last in, first out) stack.*/
package main

import "fmt"

type stack struct {
	top   int
	items [10]int
}

func (s *stack) push(item int) {
	if s.top == cap(s.items) {
		fmt.Println("stack full")
		return
	}
	s.top++
	s.items[s.top] = item
}

func (s *stack) pop() int {
	poped := s.items[s.top]
	s.items[s.top] = 0
	s.top--

	return poped
}

func main() {
	stk := stack{
		top:   0,
		items: [10]int{},
	}

	stk.push(1)
	fmt.Println(stk)
	stk.push(2)
	stk.push(4)
	stk.push(5)
	fmt.Println(stk)

	num := stk.pop()
	fmt.Println(stk, "num: ", num)
}
