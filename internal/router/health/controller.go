package health

import (
	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func NewHealthController() (*HealthController, error) {
	return &HealthController{}, nil
}

type healthStatus struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

// @Summary Health check
// @Tags Health
// @Produce  json
// @Success 200 {object} healthStatus
// @Router /health [get]
func (healthController *HealthController) Health(c *gin.Context) {
	c.JSON(200, healthStatus{Status: "Up", Error: nil})
}
