package chromium_generator

import (
	"github.com/mattanapol/kaewsai-pdf/internal/domain"
	"go.uber.org/fx"
)

var Module = fx.Module("ChromiumGenerator",
	fx.Provide(func(fileRepository domain.FileRepository) domain.PdfGenerator {
		return NewChromiumGenerator(fileRepository)
	}),
)
