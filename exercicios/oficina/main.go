package main

import "github.com/cristian-95/go-exercicios/exercicios/oficina/oficina"

func main() {

	car1 := oficina.Veiculo{Nome: "astra", Ano: 2004}
	serv := oficina.Servico{}
	serv.setServico(2)
	reg := oficina.Registro{Veiculo: car1}
	reg.addServico(serv)
	serv.setServico(1)
	reg.addServico(serv)

	car2 := oficina.Veiculo{Nome: "jacare", Ano: 1978}
	serv2 := oficina.Servico{}
	serv2.setServico(1)
	serv2.setServico(3)
	reg2 := oficina.Registro{Veiculo: car2, Servicos: serv2}

	lista := []oficina.Registro{reg, reg2}
	of := oficina.Oficina{ListaDeOrdens: lista}
	of.atender()
	of.atender()

}
