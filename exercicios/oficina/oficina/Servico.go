package oficina

type Servico struct {
	servico string
}

func (s *Servico) setServico(opt int) {
	switch opt {
	case 1:
		s.servico = "troca de oleo"
		break
	case 2:
		s.servico = "troca de pneu"
		break
	case 3:
		s.servico = "limpeza de motor"
		break
	default:
		s.servico = "invalido"
	}
}

func (s *Servico) getServico() string {
	return s.servico
}
