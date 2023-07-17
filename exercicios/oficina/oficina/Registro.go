package oficina

type Registro struct {
	Veiculo  Veiculo
	Servicos []Servico
}

func (r *Registro) getVeiculo() Veiculo {
	return r.Veiculo
}

func (r *Registro) setVeiculo(v Veiculo) {
	r.Veiculo = v
}

func (r *Registro) getServicos() []Servico {
	return r.Servicos
}

func (r *Registro) addServico(s Servico) {
	r.Servicos = append(r.Servicos, s)
}
