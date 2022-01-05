package controllers

import (
	"log"

	"github.com/aramsm/wear-store/repository"
	"github.com/gin-gonic/gin"
)

func ShowAddress(c *gin.Context) {
	street := c.Param("street")
	log.Printf("Looking for address with Street=%v", street)
	address := repository.GetAdress(street)

	if address.Street == "" {
		c.JSON(404, gin.H{
			"data": "Address not found",
		})

		return
	}

	c.JSON(200, gin.H{
		"data": address,
	})
}

func ListAddresses(c *gin.Context) {
	addresses := repository.GetAddresses()

	c.JSON(200, gin.H{
		"data": addresses,
	})
}
