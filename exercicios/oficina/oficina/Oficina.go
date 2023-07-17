package oficina

import "fmt"

type Oficina struct {
	listaDeOrdens []Registro
}

func (o *Oficina) atender() Registro {
	r := o.listaDeOrdens[0]
	o.exibir(r)
	o.listaDeOrdens = o.listaDeOrdens[1:]
	return r
}

func (o Oficina) exibir(r Registro) {
	fmt.Println("ordem de serviço:")
	fmt.Printf("Veiculo: %s %d\n", r.veiculo.getNome(), r.veiculo.getAno())
	fmt.Println("Serviços:")
	for i := 0; i < len(r.servicos); i++ {
		fmt.Printf("\t%s\n", r.servicos[i])
	}
	fmt.Println("--------------------------------")
}
