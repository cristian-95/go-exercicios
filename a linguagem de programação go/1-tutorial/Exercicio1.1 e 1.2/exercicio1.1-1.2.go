package main

import (
	"fmt"
	"os"
	"strings"
)

/*	Exercicio1.1 - Modifique o programa echo para exibir tambem os.Args[0],
que é o nome do comando que o chamou*/
func f1() {
	fmt.Println(strings.Join(os.Args[:], " "))
}

/*	Exercicio1.2 - Modifique o programa echo para exibir o indice e o
valor de cada um de seus argumentos, um por linha. */
func f2() {
	for index, value := range os.Args[:] {
		fmt.Printf("[%d] - %s\n", index, value)
	}
}

func main() {
	f1()
	f2()
}
