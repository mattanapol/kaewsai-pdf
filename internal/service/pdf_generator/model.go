package pdf_generator

import (
	"context"
	"io"

	"github.com/mattanapol/kaewsai-pdf/internal/domain"
)

type PdfGenerator interface {
	GenerateFromLink(context context.Context, url string, Options *domain.PdfGenerateRequestOption) (io.Reader, error)
	GeneratorName() string
}

type PdfGeneratorServicer interface {
	Generate(context context.Context, request *domain.PdfGenerateRequestForm) error
}

type PdfGenerateRequestReceiver interface {
	Receive(context context.Context) error
}
