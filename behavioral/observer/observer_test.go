package observer

import (
	"testing"
	"time"
)

func TestObserver(t *testing.T) {
	subscriber := NewSubscriber()
	publisher := NewPublisher()
	publisher.RegisterObserver(subscriber)

	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <-stop:
			return
		case <-tick:
			publisher.Notify()
		}
	}
}
