package adapter

import "fmt"

type Target interface {
	DoSomething()
}

type Adaptee struct {
}

func NewAdaptee() *Adaptee {
	return &Adaptee{}
}

func (a *Adaptee) Greeting() {
	fmt.Println("Greet")
}

type Adapter struct {
	Adaptee *Adaptee
}

func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{Adaptee: adaptee}
}

func (a *Adapter) DoSomething() {
	a.Adaptee.Greeting()
}
