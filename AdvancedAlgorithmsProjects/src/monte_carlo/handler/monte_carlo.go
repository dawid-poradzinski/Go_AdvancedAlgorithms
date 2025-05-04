package monte_carlo_handler

import (
	monte_carlo_logic "AdvancedAlgorithmsProjects/src/monte_carlo/logic"
	monte_carlo_model "AdvancedAlgorithmsProjects/src/monte_carlo/model"

	"github.com/gin-gonic/gin"
)

func MonteCarlo(c *gin.Context) {
	var req monte_carlo_model.MonteCarloRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if req.Iterations <= 0 {
		c.JSON(400, gin.H{"error": "Iterations must be greater than 0"})
		return
	}

	integral, percent := monte_carlo_logic.CalculateInegralValue(req.Iterations)

	c.JSON(200, monte_carlo_model.MonteCarloResponse{
		Estimate: integral,
		Percent:  percent,
	})
}
