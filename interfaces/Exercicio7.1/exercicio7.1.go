/*Exercício7.1:Usando as ideias de ByteCounter, implemente contadores para palavras e linhas.
Você perceberá que bufio.ScanWords pode ser util.
//OBS esta solução está bem 'básica' e pode ser aprimorada*/
package main

import (
	"bufio"
	"fmt"
)

func counter(s string) (int, int) {
	str := s + " " // É preciso adicionar um espaço ao final da string para que bufio.ScanWords possa scanear a ultima palavra.
	wc, lc, pos := 0, 1, 0
	for {
		advance, _, _ := bufio.ScanWords([]byte(str), false)
		str = str[advance:]
		pos += advance

		if pos < len(s) { // Verificação necessária para evitar erros do tipo: index out of range [x] with lenght [x]
			if rune(s[pos-1]) == '\n' {
				lc++
			}
		}
		if advance > 0 {
			wc++
		} else {
			break
		}
	}
	return wc, lc
}

func main() {
	spam := "exemplo de texto com\num total de quatro linhas\ne doze\npalavras"
	x, y := counter(spam)
	fmt.Println("total de palavras:", x)
	fmt.Println("total de linhas:", y)
	//	OUTPUT:
	//total de palavras: 12
	//total de linhas: 4
}
