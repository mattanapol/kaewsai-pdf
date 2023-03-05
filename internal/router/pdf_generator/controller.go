package pdf_generator

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/gin-gonic/gin"
	"github.com/mattanapol/kaewsai-pdf/internal/domain"
	"github.com/mattanapol/kaewsai-pdf/internal/service/pdf_generator"
)

type PdfGeneratorController struct {
	pdfGeneratorRequester pdf_generator.PdfGenerationRequestServicer
}

// New pdf generator controller
func NewPdfGeneratorController(pdfGeneratorRequester pdf_generator.PdfGenerationRequestServicer, sqsClient *sqs.Client) *PdfGeneratorController {
	return &PdfGeneratorController{
		pdfGeneratorRequester: pdfGeneratorRequester,
	}
}

func (controller *PdfGeneratorController) InitEndpoints(r gin.IRoutes) {
	r.POST("/generate-pdf/link", controller.GenerateFromLink)
}

type GenerateFromLinkRequest struct {
	// [wkhtmltopdf, chromium]
	Generator string `json:"generator"`
	Url       string `json:"url"`

	Options GenerateOptionRequest `json:"options"`
}

type GenerateOptionRequest struct {
	Landscape *bool    `json:"landscape"`
	Scale     *float32 `json:"scale"`
}

// Generate pdf from link
func (c *PdfGeneratorController) GenerateFromLink(context *gin.Context) {
	var generateFromLinkRequest GenerateFromLinkRequest
	if err := context.ShouldBindJSON(&generateFromLinkRequest); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Generate pdf from link
	err := c.pdfGeneratorRequester.Request(context, domain.Generator(generateFromLinkRequest.Generator),
		domain.PdfGenerateRequestForm{
			Url: generateFromLinkRequest.Url,
			Options: domain.PdfGenerateRequestOption{
				Landscape: generateFromLinkRequest.Options.Landscape,
				Scale:     generateFromLinkRequest.Options.Scale,
			},
		})

	if err != nil {
		// If error is InvalidRequestTypeError then return 400
		if _, ok := err.(*domain.InvalidRequestTypeError); ok {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		} else {
			context.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	// Return no content
	context.Status(204)
}
