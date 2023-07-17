package veiculo

type Veiculo struct {
	nome string
	ano  int
}

func (v *Veiculo) setNome(nome string) {
	v.nome = nome
}

func (v *Veiculo) getNome() string {
	return v.nome
}

func (v *Veiculo) setAno(ano int) {
	v.ano = ano
}

func (v *Veiculo) getAno() int {
	return v.ano
}
