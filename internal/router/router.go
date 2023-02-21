package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattanapol/kaewsai-pdf/internal/router/health"
)

// InitRouter initialize routing information
func InitRouter(server *http.Server,
	healthController *health.HealthController,
) {
	r := gin.New()

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", healthController.Health)

	server.Handler = r
}
