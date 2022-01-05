package server

import (
	"github.com/aramsm/wear-store/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(r *gin.Engine) *gin.Engine {
	v1 := r.Group("api/v1")
	stores := v1.Group("stores")
	addresses := v1.Group("addresses")
	employees := v1.Group("employees")

	stores.GET("/", controllers.ListStores)
	stores.GET("/:id", controllers.ShowStore)
	addresses.GET("/", controllers.ListAddresses)
	addresses.GET("/:street", controllers.ShowAddress)
	employees.GET("/", controllers.ListEmployees)
	employees.GET("/:id", controllers.ShowEmployee)

	return r
}
