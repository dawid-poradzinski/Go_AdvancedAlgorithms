package main

import (
	bisection "AdvancedAlgorithmsProjects/src/bisection/handler"
	determinant "AdvancedAlgorithmsProjects/src/determinant/handler"
	monte_carlo_handler "AdvancedAlgorithmsProjects/src/monte_carlo/handler"
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

	r.POST("/bisection", func(c *gin.Context) {
		bisection.Bisection(c)
	})

	r.POST("/monteCarlo", func(c *gin.Context) {
		monte_carlo_handler.MonteCarlo(c)
	})

	r.Run(":8080")

}
