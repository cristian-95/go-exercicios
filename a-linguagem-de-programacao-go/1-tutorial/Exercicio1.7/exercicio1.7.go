// Exercicio 1.7 - A chamada da função io.Copy(dst, src) lê de src e escreve em dst.
// Use-a no lugar de io.ReadAll para copiar o corpo da resposta
// para os.Stdout sem exigir um buffer grande o suficiente para armazenar todo o stream.
// Não se esqueça de verificar o resultado do erro de io.Copy.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("\n\n%v bytes copied from %s\n", b, url)
	}

}
