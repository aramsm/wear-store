package controllers

import (
	"log"

	"github.com/aramsm/wear-store/repository"
	"github.com/gin-gonic/gin"
)

func ShowStore(c *gin.Context) {
	data := repository.ReadAndParseData()
	id := c.Param("id")
	log.Printf("Looking for store with id=%v", id)
	store := repository.GetStore(data, id)

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

func ListStoresByBrand(c *gin.Context) {
	data := repository.ReadAndParseData()
	brand := c.Param("brand_label")
	log.Printf("Looking for store with brand_label=%v", brand)

	stores := repository.GetStoresByBrand(data, brand)

	c.JSON(200, gin.H{
		"data": stores,
	})
}
