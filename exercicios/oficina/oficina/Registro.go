package oficina

import (
	"github.com/cristian-95/go-exercicios/exercicios/oficina/veiculo"
)

type Registro struct {
	Veiculo  veiculo.Veiculo
	Servicos []Servico
}

func (r *Registro) GetVeiculo() veiculo.Veiculo {
	return r.Veiculo
}

func (r *Registro) SetVeiculo(v veiculo.Veiculo) {
	r.Veiculo = v
}

func (r *Registro) GetServicos() []Servico {
	return r.Servicos
}

func (r *Registro) AddServico(s Servico) {
	r.Servicos = append(r.Servicos, s)
}
