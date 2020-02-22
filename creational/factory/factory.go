package factory

import "errors"

type PizzaType int

const (
	ChicagoType PizzaType = 1 << iota
	GreekType
	CaliforniaType
)

type PizzaStore struct {
}

func NewPizzaStore() *PizzaStore {
	return &PizzaStore{}
}

func (p *PizzaStore) createPizza(pizzaType PizzaType) (Pizza, error) {
	switch pizzaType {
	case ChicagoType:
		return NewChicagoPizza(), nil
	case GreekType:
		return NewGeekPizza(), nil
	case CaliforniaType:
		return NewChicagoPizza(), nil
	default:
		return nil, errors.New("not found pizza type")
	}
}

type Pizza interface {
	Get() string
}

type ChicagoPizza struct {
}

func NewChicagoPizza() *ChicagoPizza {
	return &ChicagoPizza{}
}

func (c *ChicagoPizza) Get() string {
	return "chicago"
}

type GeekPizza struct {
}

func NewGeekPizza() *GeekPizza {
	return &GeekPizza{}
}

func (g *GeekPizza) Get() string {
	return "greek"
}

type CaliforniaPizza struct {
}

func NewCaliforniaPizza() *CaliforniaPizza {
	return &CaliforniaPizza{}
}

func (c *CaliforniaPizza) Get() string {
	return "california"
}
