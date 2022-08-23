/*Exercicio 4.9
Escreva um programa wordfreq para informar a frequencia de cada
palavra em um arquivo-texto de entrada. Chame input.Split(bufio.ScanWords)
antes da primeira chamada a Scan para separar a entrada em palavras,
e não em linhas

TO DO:
	implementar usando bufio.ScanWords
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func wordFreq(text string) (words map[string]int) {
	t := strings.Split(text, " ")
	for _, v := range t {
		words[string(v)]++
	}
	return
}

func main() {
	fileName := os.Args[1]
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	frequencias := wordFreq(string(fileContent))
	for word, freq := range frequencias {
		fmt.Printf("%v: %v\n", word, freq)
	}
}
