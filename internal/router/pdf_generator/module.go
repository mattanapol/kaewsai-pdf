package pdf_generator

import (
	"github.com/mattanapol/kaewsai-pdf/internal/service/pdf_generator"
	"go.uber.org/fx"
)

var Module = fx.Module("PdfGeneratorControllers",
	pdf_generator.RequestModule,
	fx.Provide(NewPdfGeneratorController),
)
