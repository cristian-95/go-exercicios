package oficina

type Veiculo struct {
	Nome string
	Ano  int
}

func (v *Veiculo) SetNome(nome string) {
	v.Nome = nome
}

func (v *Veiculo) GetNome() string {
	return v.Nome
}

func (v *Veiculo) SetAno(ano int) {
	v.Ano = ano
}

func (v *Veiculo) GetAno() int {
	return v.Ano
}
