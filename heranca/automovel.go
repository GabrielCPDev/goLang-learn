package heranca

type Automovel struct {
	Ano int
	Placa string
	Modelo string
}

type Moto struct {
	Automovel
	Cilindradas int
}

type Carro struct {
	Automovel
	QttPortas int
	Potencia int
	possuiArCondicionado bool
}