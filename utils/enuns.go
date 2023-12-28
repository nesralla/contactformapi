package utils

type CarroTipo int
type Modalidade int

const (
	Carros CarroTipo = iota + 1
	Motos
	Caminhoes
)
const (
	Mobilidade Modalidade = iota
	Transporte
	Delivery
	Outros
)

func (d CarroTipo) String() string {
	switch d {
	case Carros:
		return "Carros"
	case Motos:
		return "Motos"
	case Caminhoes:
		return "Caminhoes"

	default:
		return "Invalid tipo"
	}
}

func (d Modalidade) String() string {
	switch d {
	case Mobilidade:
		return "Mobilidade"
	case Transporte:
		return "Transporte"
	case Delivery:
		return "Outros"

	default:
		return "Invalid tipo"
	}
}
