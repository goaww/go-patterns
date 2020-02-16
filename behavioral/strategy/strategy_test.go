package strategy

import "testing"

func TestNewDuck(t *testing.T) {
	duck := NewDuck(NewFlyableImpl(), NewQuack())
	duck.PerformFly()
	duck.PerformQuack()
}
