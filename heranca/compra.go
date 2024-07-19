package heranca

import "time"

type Compra struct {
	Mercado string
	Data    time.Time
	Items   []ItemCompra
}

type ItemCompra struct {
	Nome string
}

 func NewCompra(mercado string, data time.Time, nomeDoitens [] string) *Compra {
	var itens []ItemCompra

	for _ , nome := range nomeDoitens {
		itens = append(itens, ItemCompra{Nome: nome})
	}
	return &Compra { Mercado: mercado, Data: data, Items: itens}
 }