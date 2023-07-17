package oficina

type Carro struct {
	Veiculo
	nPassageiros int
}

func (c *Carro) getNPassageiros() int {
	return c.nPassageiros
}

func (c *Carro) setNPassageiros(num int) {
	c.nPassageiros = num
}
