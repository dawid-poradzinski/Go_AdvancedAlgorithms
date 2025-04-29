package main

import (
	determinant "AdvancedAlgorithmsProjects/src/determinant/handler"
	quadruplet "AdvancedAlgorithmsProjects/src/quadruple/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/quadruplet/", func(c *gin.Context) {
		quadruplet.PrimeQuadruplet(c)
	})

	r.POST("/determinant", func(c *gin.Context) {
		determinant.Determinant(c)
	})

	r.Run(":8080")

}
