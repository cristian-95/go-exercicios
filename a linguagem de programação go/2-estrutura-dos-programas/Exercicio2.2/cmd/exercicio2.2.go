package main

import (
	"fmt"
	"os"
	"strconv"

	"Exercicio2.2/massunits"
)

func main() {
	for _, arg := range os.Args[1:] {
		m, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex2.2: %v \n", err)
			os.Exit(1)
		}
		g := massunits.Grama(m)
		kg := massunits.Quilograma(m)
		lb := massunits.Libra(m)
		oz := massunits.Onça(m)

		fmt.Printf("%s = %s = %s = %s\n",
			g, massunits.GToKg(g), massunits.GToLb(g), massunits.GToOz(g))
		fmt.Printf("%s = %s = %s = %s\n",
			kg, massunits.KgToG(kg), massunits.KgToLb(kg), massunits.KgToOz(kg))
		fmt.Printf("%s = %s = %s = %s\n",
			lb, massunits.LbToG(lb), massunits.LbToKg(lb), massunits.LbToOz(lb))
		fmt.Printf("%s = %s = %s = %s\n",
			oz, massunits.OzToG(g), massunits.OzToKg(oz), massunits.OzToLb(oz))
	}

}
