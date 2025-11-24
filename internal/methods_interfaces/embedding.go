package methods_interfaces

import "fmt"

// Logger is a simple embedded type example.
type Logger struct{}

func (Logger) Log(msg string) {
	fmt.Println("[LOG]", msg)
}

type Server struct {
	Logger
	Addr string
}

func EmbeddingDemo() {
	s := Server{Addr: ":8080"}
	s.Log("server starting at " + s.Addr)
}
