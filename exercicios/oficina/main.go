package main

import (
	"github.com/cristian-95/go-exercicios/exercicios/oficina/oficina"
	"github.com/cristian-95/go-exercicios/exercicios/oficina/veiculo"
)

func main() {

	car1 := veiculo.Veiculo{Nome: "astra", Ano: 2004}
	serv := oficina.Servico{}
	serv.SetServico(2)
	reg := oficina.Registro{Veiculo: car1}
	reg.AddServico(serv)
	serv.SetServico(1)
	reg.AddServico(serv)

	car2 := veiculo.Veiculo{Nome: "jacare", Ano: 1978}
	reg2 := oficina.Registro{Veiculo: car2}
	serv.SetServico(3)
	reg2.AddServico(serv)
	serv.SetServico(2)
	reg2.AddServico(serv)
	serv.SetServico(1)
	reg2.AddServico(serv)

	lista := []oficina.Registro{reg, reg2}
	of := oficina.Oficina{ListaDeOrdens: lista}
	of.Atender()
	of.Atender()

}
