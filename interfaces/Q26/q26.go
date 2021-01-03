/*Q26. (2) Interfaces and max()

In exercise Q13 we created a max function that works on a slice of integers. The
question now is to create a program that shows the maximum number and that
works for both integers and floats. Try to make your program as generic as possible,
although that is quite difficult in this case.
*/
package main

import (
	"fmt"
)

func main() {
	x := max(2, 3, 5, 10)
	y := max(12.4, 3.35, 5.9, 10.05)
	fmt.Println("max from {2, 3, 5, 10}:", x, "\nmax from {12.4, 3.35, 5.9, 10.05}:", y)
}

func max(sl ...interface{}) interface{} {
	switch t := sl[0].(type) {
	case int:
		result := 0
		for i := 0; i < len(sl); i++ {
			if sl[i].(int) > result {
				result = sl[i].(int)
			}
		}
		return result
	case float64:
		result := 0.0
		for i := 0; i < len(sl); i++ {
			if sl[i].(float64) > result {
				result = sl[i].(float64)
			}
		}
		return result
	default:
		fmt.Println("ERROR: Wrong type.\nExpect: int or float64, got: ", t)
		return false
	}
}
