package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sja-boiler-plate/core/models"
)

//
//func getProduct(c *gin.Context) {
//	code := c.Param("code")
//
//	product := models.GetProduct(code)
//
//	if product == nil {
//		c.AbortWithStatus(http.StatusNotFound)
//	} else {
//		c.IndentedJSON(http.StatusOK, product)
//	}
//}

func getProducts(c *gin.Context) {
	product, err := models.GetProducts()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, product)
	}
}
