/*
Exercicio 4.3 Reescreva a função reverse usando um ponteiro de array no lugar da fatia
*/
package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	reverse(&array)
	fmt.Println(array)
}

func reverse(input *[5]int) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}
