package main

import (
	"fmt"
	"os"
	"strconv"

	conv "github.com/cristian-95/go-exercicios/a-linguagem-de-programação-go/2-estrutura-dos-programas/Exercicio2.2/massunits"
)

func main() {
	for _, arg := range os.Args[1:] {
		m, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex2.2: %v \n", err)
			os.Exit(1)
		}
		g := conv.Grama(m)
		kg := conv.Quilograma(m)
		lb := conv.Libra(m)
		oz := conv.Onça(m)

		fmt.Printf("%s = %s = %s = %s\n",
			g, conv.GToKg(g), conv.GToLb(g), conv.GToOz(g))
		fmt.Printf("%s = %s = %s = %s\n",
			kg, conv.KgToG(kg), conv.KgToLb(kg), conv.KgToOz(kg))
		fmt.Printf("%s = %s = %s = %s\n",
			lb, conv.LbToG(lb), conv.LbToKg(lb), conv.LbToOz(lb))
		fmt.Printf("%s = %s = %s = %s\n",
			oz, conv.OzToG(oz), conv.OzToKg(oz), conv.OzToLb(oz))
	}

}
