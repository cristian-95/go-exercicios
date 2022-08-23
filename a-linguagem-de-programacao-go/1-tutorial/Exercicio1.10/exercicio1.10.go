/* Exercicio 1.10 - Encontre um site que gere uma quantidade grande de dados.
 * Investigue o caching executando fetchall duas vezes sucessivamente para ver
 * se o tempo informado sofre muita alteração. Você sempre obtem o mesmo
 * conteúdo? Modifique fetchall para exibir sua saída em um arquivo para que ela
 * possa ser examinada.
 */
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	filePointer, err := os.Create("output.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating file: %s", err)
		os.Exit(1)
	}
	defer filePointer.Close()

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Fprintf(filePointer, "%s\n", <-ch)
	}

	fmt.Fprintf(filePointer, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
