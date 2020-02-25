package command

import "testing"

func TestHelloCommand(t *testing.T) {
	receiver := NewHelloReceiver()
	command := NewHelloCommand(receiver)

	invoker := NewInvoker()
	invoker.SetCommand(command)

	invoker.Invoke()

}
