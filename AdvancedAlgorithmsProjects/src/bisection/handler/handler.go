package bisection

import (
	bisection_model "AdvancedAlgorithmsProjects/src/bisection/model"
	bisection_logic "AdvancedAlgorithmsProjects/src/bisection/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Bisection(c *gin.Context) {
	var req bisection_model.BisectionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := bisection_logic.ValidateRange(req.Delta, req.Min, req.Max)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := bisection_logic.Bisection(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response := bisection_model.BisectionResponse{
		X: result,
	}

	c.JSON(http.StatusOK, response)
}
