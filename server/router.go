package server

import "github.com/gin-gonic/gin"

// Loads the routes configurations
func ConfigRoutes(r *gin.Engine) *gin.Engine {
	v1 := r.Group("api/v1")
	stores := v1.Group("stores")
	addresses := v1.Group("addresses")
	employees := v1.Group("employees")

	stores.GET("/")
	addresses.GET("/")
	employees.GET("/")

	return r
}
