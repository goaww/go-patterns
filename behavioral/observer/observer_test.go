package observer

import "testing"

func TestPublisher_RegisterObserver(t *testing.T) {
	subscriber := NewSubscriber()
	publisher := NewPublisher()
	publisher.RegisterObserver(subscriber)

	publisher.Notify()
}
