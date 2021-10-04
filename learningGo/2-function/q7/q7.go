/*Q7. (0) Integer ordering
1. Write a function that returns its (two) parameters in the right, numerical (ascend-
ing) order:
	f(7,2) → 2,7
	f(2,7) → 2,7 */
package main

import "fmt"

func order(x, y int) (int, int) {
	if x <= y {
		return x, y
	} else {
		return y, x
	}
}

func main() {
	fmt.Println(order(82, 8))
}
