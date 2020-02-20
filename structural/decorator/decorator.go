package decorator

import "fmt"

//The Decorator Pattern attaches additional responsibilities to an object dynamically.
//Decorators provide a flexible alternative to subclassing for extending functionality.

type Logger interface {
	Info(interface{}) string
}
type DefaultLogger struct {
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

func (d *DefaultLogger) Info(arg interface{}) string {
	return fmt.Sprint("%i", arg)
}

type SimpleLogger struct {
	Log Logger
}

func NewSimpleLogger(log Logger) *SimpleLogger {
	return &SimpleLogger{Log: log}
}

func (s *SimpleLogger) Info(arg interface{}) string {
	i := s.Log.Info(arg)
	rs := fmt.Sprintf("Decorator %s", i)
	return rs
}
