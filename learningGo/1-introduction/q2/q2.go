/*Q2. (0) For-loop
	1. Create a simple loop with the for construct. Make it loop 10 times and print out
the loop counter with the fmt package.
	2. Rewrite the loop from 1. to use goto. The keyword for may not be used.
	3. Rewrite the loop again so that it fills an array and then prints that array to the
screen. */

package main

import "fmt"

func main() {

	//1.
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//2.
	// 	i := 0
	// LOOP:
	// 	fmt.Println(i)
	// 	i++
	// 	if i < 10 {
	// 		goto LOOP
	// 	}

	//3.
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)

}
