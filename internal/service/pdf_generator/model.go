package pdf_generator

import (
	"context"

	"github.com/mattanapol/kaewsai-pdf/internal/domain"
)

type PdfGeneratorServicer interface {
	Generate(context context.Context, request *domain.PdfGenerateRequestForm) error
}
