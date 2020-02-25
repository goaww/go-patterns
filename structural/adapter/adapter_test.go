package adapter

import "testing"

func TestAdapter(t *testing.T) {
	var target Target
	target = NewAdapter(NewAdaptee())
	target.DoSomething()
}
