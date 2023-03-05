package pdf_generator

import (
	"github.com/mattanapol/kaewsai-pdf/internal/service/pdf_generator/requester"
	"go.uber.org/fx"
)

var Module = fx.Module("PdfGeneratorService",
	fx.Provide(NewPdfGeneratorService),
)

var RequestModule = fx.Module("PdfGeneratorRequest",
	requester.Module,
	fx.Provide(NewPdfGenerationRequestService),
)
