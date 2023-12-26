package utils

type CarroTipo int
type Modalidade int

const (
	Carros CarroTipo = iota
	Moto
	Caminhao
)
const (
	Mobilidade Modalidade = iota
	Transporte
	Delivery
	Outros
)
