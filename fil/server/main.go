package main

import (
	"context"
	"log"
)

func main() {
	s := Server{
		Address: "localhost:62715",
		Logger:  log.Default(),
		Exit:    make(chan struct{}),
	}

	err := s.Run(context.Background())
	if err != nil {
		panic(err)
	}

	s.Wait()
}
