package decorator

import "testing"

func TestNewSimpleLogger(t *testing.T) {
	logger := NewSimpleLogger(NewDefaultLogger())
	logger.Info("Say hello")
}
