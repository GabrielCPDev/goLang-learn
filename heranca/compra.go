package heranca

import (
	"errors"
	"time"
)

type Compra struct {
	Mercado string
	Data    time.Time
	Items   []ItemCompra
}

type ItemCompra struct {
	Nome string
}

 func NewCompra(mercado string, data time.Time, nomeDoitens [] string) (*Compra, error) {
	var itens []ItemCompra


	for _ , nome := range nomeDoitens {
		itens = append(itens, ItemCompra{Nome: nome})
	}

	if mercado == "" || len(nomeDoitens) == 0 {
		return nil, errors.New("mercado é obrigatório!")
	}
	return &Compra { Mercado: mercado, Data: data, Items: itens}, nil
 }