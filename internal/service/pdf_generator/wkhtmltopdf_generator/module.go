package wkhtmltopdf_generator

import (
	"github.com/mattanapol/kaewsai-pdf/internal/persistence/file_repos"
	"github.com/mattanapol/kaewsai-pdf/internal/service/pdf_generator"
	"go.uber.org/fx"
)

var Module = fx.Module("WkhtmltopdfGenerator",
	fx.Provide(func(fileRepository file_repos.FileRepository) pdf_generator.PdfGenerator {
		return NewWkhtmltopdfGenerator(fileRepository)
	}),
)
