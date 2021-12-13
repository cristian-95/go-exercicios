/*(0) Stack as package
1. See the Q9 exercise. In this exercise we want to create a  separate package for that code. Create a proper package for your stack implementation, Push, Pop and the Stack type need to be exported */
package main

import (
	"fmt"

	"q16.go/mystack"
)

func main() {
	mystk := mystack.New()

	mystk.Push(2)
	fmt.Println(mystk)

	mystk.Push(1)
	mystk.Push(3)
	mystk.Push(5)
	mystk.Push(4)
	mystk.Push(7)
	fmt.Println(mystk)

	num := mystk.Pop()
	fmt.Println(mystk, "poped num: ", num)
}
