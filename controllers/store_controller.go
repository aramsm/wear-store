package controllers

import (
	"log"

	"github.com/aramsm/wear-store/repository"
	"github.com/gin-gonic/gin"
)

func ShowStore(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Looking for store with id=%v", id)
	store := repository.GetStore(id)

	if store.ID == "" {
		c.JSON(404, gin.H{
			"data": "Store not found",
		})

		return
	}

	c.JSON(200, gin.H{
		"data": store,
	})
}

func ListStores(c *gin.Context) {
	stores := repository.ReadAndParseData()

	c.JSON(200, gin.H{
		"data": stores,
	})
}
