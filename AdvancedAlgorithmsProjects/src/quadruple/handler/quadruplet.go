package quadruplet

import (
	"AdvancedAlgorithmsProjects/src/quadruple/model"
	"AdvancedAlgorithmsProjects/src/quadruple/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PrimeQuadruplet(c *gin.Context) {
	var req model.QuadrupletRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json"})
		return
	}

	if req.End < 2 || req.Start > req.End {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid range: Start must be <= End and End >= 2"})
		return
	}

	result := service.AnalyzeQuadruplets(req.Start, req.End)

	c.JSON(http.StatusOK, result)
}
