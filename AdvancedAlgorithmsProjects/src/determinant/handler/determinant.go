package determinant

import (
	"AdvancedAlgorithmsProjects/src/determinant/model"
	"AdvancedAlgorithmsProjects/src/determinant/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Determinant(c *gin.Context) {
	var req model.DeterminantRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !service.CheckIfValidMatrix(req.Size, req.Matrix) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size to matrix ratio"})
		return
	}

	result, err := service.CalculateDeterminant(req.Size, req.Matrix)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	response := model.DeterminantResponse{
		Determinant: result,
	}

	c.JSON(http.StatusOK, response)
}
