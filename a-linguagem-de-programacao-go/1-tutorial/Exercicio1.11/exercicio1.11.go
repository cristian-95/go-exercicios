/* Exercicio 1.11 - Experimente usar fetchall com listas mais longas de argumentos a...
 * [...]
 * Como o programa se comporta dse um site simplesmente nao responder?
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
	outputPointer, err := os.Create("output.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating file: %s", err)
		os.Exit(1)
	}

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Fprintf(outputPointer, "%s\n", <-ch)
	}

	fmt.Fprintf(outputPointer, "%.2fs elapsed\n", time.Since(start).Seconds())
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
