package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Engine *gin.Engine
}

// Returns a server. If the port argument is blank,
// sets the server to `3000` port
func NewServer(port string) Server {
	if port == "" {
		return Server{Port: "3000", Engine: gin.Default()}
	}

	return Server{Port: port, Engine: gin.Default()}
}

// Loads the routes configurations and then runs the server
func (s *Server) Run() {
	r := ConfigRoutes(s.Engine)

	log.Printf("Server running at port: %v", s.Port)
	log.Fatal(r.Run(":" + s.Port))
}
