package pdf_generator

import (
	"go.uber.org/fx"
)

var Module = fx.Module("PdfGeneratorService",
	fx.Provide(NewPdfGeneratorService),
)
