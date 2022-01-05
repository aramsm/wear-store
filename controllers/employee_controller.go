package controllers

import (
	"log"

	"github.com/aramsm/wear-store/repository"
	"github.com/gin-gonic/gin"
)

func ShowEmployee(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Looking for address with id=%v", id)
	employee := repository.GetEmployee(id)

	if employee.ID == "" {
		c.JSON(404, gin.H{
			"data": "Employee not found",
		})

		return
	}

	c.JSON(200, gin.H{
		"data": employee,
	})
}

func ListEmployees(c *gin.Context) {
	employees := repository.GetEmployees()

	c.JSON(200, gin.H{
		"data": employees,
	})
}
