package controllers

import (
	"github.com/aramsm/wear-store/repository"
	"github.com/gin-gonic/gin"
)

func ShowStore(c *gin.Context) {
	data := repository.ReadAndParseData()

	c.JSON(200, gin.H{
		"data": data,
	})
}
