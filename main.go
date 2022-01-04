package main

import (
	"fmt"

	"github.com/aramsm/wear-store/server"
)

func main() {
	s := server.NewServer("3000")
	s.Run()
	fmt.Printf("Server up at %v", s.Port)
}
