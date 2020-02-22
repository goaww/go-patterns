package factory

import (
	"fmt"
	"testing"
)

func TestPizza(t *testing.T) {
	store := NewPizzaStore()
	pizza, _ := store.createPizza(CaliforniaType)
	fmt.Println(pizza.Get())

}
