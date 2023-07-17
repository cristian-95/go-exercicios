package oficina

type Caminhao struct {
	Veiculo
	nEixos int
}

func (c *Caminhao) getnEixos() int {
	return c.nEixos
}

func (c *Caminhao) setnEixos(num int) {
	c.nEixos = num
}
