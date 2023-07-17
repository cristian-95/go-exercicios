package oficina

type Registro struct {
	Veiculo  Veiculo
	Servicos []Servico
}

func (r *Registro) GetVeiculo() Veiculo {
	return r.Veiculo
}

func (r *Registro) SetVeiculo(v Veiculo) {
	r.Veiculo = v
}

func (r *Registro) GetServicos() []Servico {
	return r.Servicos
}

func (r *Registro) AddServico(s Servico) {
	r.Servicos = append(r.Servicos, s)
}
