package command

import "fmt"

type Command interface {
	Execute()
}

type Receiver interface {
	Handle()
}

type HelloReceiver struct {
}

func NewHelloReceiver() *HelloReceiver {
	return &HelloReceiver{}
}

func (h *HelloReceiver) Handle() {
	fmt.Println("Say hello")
}

type HelloCommand struct {
	Rec Receiver
}

func (h *HelloCommand) Execute() {
	h.Rec.Handle()
}

func NewHelloCommand(rec Receiver) *HelloCommand {
	return &HelloCommand{Rec: rec}
}

type Invoker struct {
	Commands []Command
}

func NewInvoker() *Invoker {
	return &Invoker{}
}

func (i *Invoker) SetCommand(com Command) {
	i.Commands = append(i.Commands, com)
}

func (i *Invoker) Invoke() {
	for _, command := range i.Commands {
		command.Execute()
	}
}
