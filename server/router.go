package server

import (
	"github.com/aramsm/wear-store/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(r *gin.Engine) *gin.Engine {
	v1 := r.Group("api/v1")
	stores := v1.Group("stores", controllers.ShowStore)
	addresses := v1.Group("addresses")
	employees := v1.Group("employees")

	stores.GET("/")
	addresses.GET("/")
	employees.GET("/")

	return r
}
