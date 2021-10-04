/*Q3. (0) FizzBuzz
1. Solve this problem, called the Fizz-Buzz [23] problem:

Write a program that prints the numbers from 1 to 100. But for multiples
of three print “Fizz” instead of the number and for the multiples of five
print “Buzz”. For numbers which are multiples of both three and five print
“FizzBuzz”.
*/
package main

import "fmt"

func main() {
	//1.
	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
