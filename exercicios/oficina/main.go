package main

import "github.com/cristian-95/go-exercicios/exercicios/oficina/oficina"

func main() {

	car1 := oficina.Carro{oficina.Veiculo{"astra", 2004}, 4}
	serv := oficina.Servico{}
	serv.setServico(2)
	serv.setServico(1)
	serv.setServico(3)
	reg := oficina.Registro{car1, serv}

	car2 := oficina.Caminhao{oficina.Veiculo{"jacare", 1978}, 3}
	serv2 := oficina.Servico{}
	serv2.setServico(1)
	serv2.setServico(3)
	reg2 := oficina.Registro{car2, serv2}

	lista := []oficina.Registro{reg, reg2}
	of := oficina.Oficina{lista}
	of.atender()
	of.atender()

}
