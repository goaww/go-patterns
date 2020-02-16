package strategy

import "fmt"

type Flyable interface {
	Fly()
}

type FlyableImpl struct {
}

func NewFlyableImpl() *FlyableImpl {
	return &FlyableImpl{}
}

func (f *FlyableImpl) Fly() {
	fmt.Println("Fly....")
}

type Quackable interface {
	Quack()
}

type Quack struct {
}

func NewQuack() *Quack {
	return &Quack{}
}

func (q *Quack) Quack() {
	fmt.Println("quack quack quack...")
}

type Duck interface {
	PerformFly()
	PerformQuack()
}
type DuckImpl struct {
	Fly   Flyable
	Quack Quackable
}

func (d *DuckImpl) PerformFly() {
	d.Fly.Fly()
}

func (d *DuckImpl) PerformQuack() {
	d.Quack.Quack()
}

func NewDuck(fly Flyable, quack Quackable) *DuckImpl {
	return &DuckImpl{Fly: fly, Quack: quack}
}
