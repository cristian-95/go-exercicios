package oficina

import (
	"fmt"
)

type Oficina struct {
	ListaDeOrdens []Registro
}

func (o *Oficina) Atender() Registro {
	r := o.ListaDeOrdens[0]
	o.Exibir(r)
	o.ListaDeOrdens = o.ListaDeOrdens[1:]
	return r
}

func (o Oficina) Exibir(r Registro) {
	fmt.Println("ordem de serviço:")
	fmt.Printf("Veiculo: %s %d\n", r.Veiculo.GetNome(), r.Veiculo.GetAno())
	fmt.Println("Serviços:")
	for i := 0; i < len(r.Servicos); i++ {
		fmt.Printf("\t%s\n", r.Servicos[i])
	}
	fmt.Println("--------------------------------")
}
