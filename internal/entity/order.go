package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.EhValido()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) EhValido() error {
	if o.ID == "" {
		return errors.New("id incorreto")
	}
	if o.Price <= 0 {
		return errors.New("price incorreto")
	}
	if o.Tax <= 0 {
		return errors.New("taxa incorreta")
	}
	return nil
}

func (o *Order) CalculaterPrecoFinal() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.EhValido()
	if err != nil {
		return err
	}
	return nil
}
