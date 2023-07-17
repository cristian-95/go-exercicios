package oficina

type Servico struct {
	Servico string
}

func (s *Servico) setServico(opt int) {
	switch opt {
	case 1:
		s.Servico = "troca de oleo"
		break
	case 2:
		s.Servico = "troca de pneu"
		break
	case 3:
		s.Servico = "limpeza de motor"
		break
	default:
		s.Servico = "invalido"
	}
}

func (s *Servico) getServico() string {
	return s.Servico
}
