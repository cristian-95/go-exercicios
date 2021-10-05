/*10. (1) Var args
Write a function that takes a variable number of ints and prints each integer on a
separate line	*/
package main

import "fmt"

func f1(numbers ...int) {
	for _, n := range numbers {
		fmt.Println(n)
	}

}

func main() {
	f1(10, 9, 8, 7, 6, 5, 4, 3, 2, 1)

}
