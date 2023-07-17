package oficina

type Veiculo struct {
	Nome string
	Ano  int
}

func (v *Veiculo) setNome(nome string) {
	v.Nome = nome
}

func (v *Veiculo) getNome() string {
	return v.Nome
}

func (v *Veiculo) setAno(ano int) {
	v.Ano = ano
}

func (v *Veiculo) getAno() int {
	return v.Ano
}
