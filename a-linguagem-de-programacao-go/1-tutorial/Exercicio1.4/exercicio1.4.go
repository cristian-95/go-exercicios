/*Exercicio 1.4
Modifique dup2 paraque exiba os nomes de todos os arquivosem que cada linha duplicada ocorre.
	dup2 pode ser encontrado na seção 1.3 do capitulo 1
	link-quebra-galho: http://www.gopl.io/ch1.pdf
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//estrategia, foi criado um novo mapa para armazenar apenas o nome dos arquivos
	filenames := make(map[string]string, len(os.Args[:]))
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("[from: %s]\t%d\t%s\t\n", filenames[line], n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int, filenames map[string]string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
		filenames[input.Text()] = f.Name() // f.Name() retorna o nome do arquivo em questão
	}
}
