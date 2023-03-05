package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattanapol/kaewsai-pdf/internal/router/health"
	"github.com/mattanapol/kaewsai-pdf/internal/router/pdf_generator"
)

// InitRouter initialize routing information
func InitRouter(server *http.Server,
	healthController *health.HealthController,
	pdfGeneratorController *pdf_generator.PdfGeneratorController,
) {
	r := gin.New()

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", healthController.Health)

	apiV1Group := r.Group("/api/v1")
	pdfGeneratorController.InitEndpoints(apiV1Group)

	server.Handler = r
}
