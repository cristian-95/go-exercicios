/*
1.3 Experimente medir a diferença de tempo entre as versões
anteriores de echo e a versão que usa strings.Join
*/
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func funcWithoutJoin() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	seconds := time.Since(start).Seconds()
	fmt.Printf("Função sem o uso de strings.Join\ntempo de execução: %v segundos\n", seconds)

}

func funcWithJoin() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	seconds := time.Since(start).Seconds()
	fmt.Printf("Função com o uso de strings.Join\ntempo de execução: %v segundos\n", seconds)
}

func main() {
	funcWithoutJoin()
	fmt.Println()
	funcWithJoin()
}
