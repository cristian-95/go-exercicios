package oficina

type Registro struct {
	veiculo  Veiculo
	servicos []Servico
}

func (r *Registro) getVeiculo() Veiculo {
	return r.veiculo
}

func (r *Registro) setVeiculo(v Veiculo) {
	r.veiculo = v
}

func (r *Registro) getServicos() []Servico {
	return r.servicos
}

func (r *Registro) addServico(s Servico) {
	r.servicos = append(r.servicos, s)
}
