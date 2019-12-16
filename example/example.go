package example

import "fmt"

// Envelope Struct
type Envelope struct {
	Payload string
}

// Message Interface
type Message interface {
	print() string
}

func (env Envelope) print() string {
	return fmt.Println(*env.Payload)
}
