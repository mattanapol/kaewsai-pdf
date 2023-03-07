package wkhtmltopdf_generator

import (
	"github.com/mattanapol/kaewsai-pdf/internal/domain"
	"go.uber.org/fx"
)

var Module = fx.Module("WkhtmltopdfGenerator",
	fx.Provide(func(fileRepository domain.FileRepository) domain.PdfGenerator {
		return NewWkhtmltopdfGenerator(fileRepository)
	}),
)
