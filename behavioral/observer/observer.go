package observer

import "fmt"

type Observable interface {
	Notify()
}

type Publisher struct {
	Observers []Observer
}

func NewPublisher() *Publisher {
	return &Publisher{}
}

func (p *Publisher) Notify() {
	for _, o := range p.Observers {
		go o.Update()
	}
}

func (p *Publisher) RegisterObserver(observer Observer) {
	p.Observers = append(p.Observers, observer)
}

type Observer interface {
	Update()
}

type Subscriber struct {
}

func NewSubscriber() *Subscriber {
	return &Subscriber{}
}

func (s *Subscriber) Update() {
	fmt.Println("updated")
}
