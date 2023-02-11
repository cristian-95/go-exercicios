/* Exercicio 4.4 - Escreva uma versão de  rotate que funcione com um unico passo */
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	for range a {
		rotate(&a)
		fmt.Println(a)
	}
}

func rotate(input *[]int) {
	aux := *input
	*input = append(aux[len(aux)-1:], aux[0:len(aux)-1]...)
}
