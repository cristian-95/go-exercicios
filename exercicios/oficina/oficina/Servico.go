package oficina

type Servico struct {
	Servico string
}

func (s *Servico) SetServico(opt int) {
	switch opt {
	case 1:
		s.Servico = "troca de oleo"
	case 2:
		s.Servico = "troca de pneu"
	case 3:
		s.Servico = "limpeza de motor"
	default:
		s.Servico = "invalido"
	}
}

func (s *Servico) GetServico() string {
	return s.Servico
}
