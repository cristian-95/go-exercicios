/*	Q11. (1) Fibonacci
1. The Fibonacci sequence starts as follows:
	1, 1, 2, 3, 5, 8, 13, . . .
Or in mathematical terms:
	 x1 = 1; x2 = 1; xn= xn−1 + xn−2 ∀n > 2.
Write a function that takes an int value and gives that any terms of the Fibonacci sequence.

📌Correção a sequencia de fibonacci começa assim:
		0, 1, 1, 2, 3, 5, 8, 13 ...
		1º = 1,
		2º = 1,
		3º = 2 ...		*/

package main

import "fmt"

//1.
func fibonacci(N int) []int {
	current := 0
	prev1 := 1
	prev2 := 0
	var sequence []int
	sequence = append(sequence, 0) // nao consegui implementar a logica deste zero no inicio :/

	for i := 0; i < N; i++ {
		sequence = append(sequence, prev1)
		current = prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return sequence
}

func main() {
	sequence := fibonacci(7)
	fmt.Println(sequence)
	// prints : [0 1 1 2 3 5 8 13]
}
