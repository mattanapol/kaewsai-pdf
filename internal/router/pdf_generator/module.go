package pdf_generator

import (
	"github.com/mattanapol/kaewsai-pdf/internal/service/pdf_generator/requester"
	"go.uber.org/fx"
)

var Module = fx.Module("PdfGeneratorControllers",
	requester.Module,
	fx.Provide(NewPdfGeneratorController),
)
