// Exercicio 1.9 - Modifique fetch para exibir também o código
// de status HTTP encontrado em resp.Status.
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

		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		stats := resp.Status
		fmt.Printf("Request to %s\nStatus: %s\n--------------------------------------------------------------------\nContent:\n\n", url, stats)
		fmt.Printf("%s", b)
	}

}
