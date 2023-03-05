package setting

import "go.uber.org/fx"

var ApiModule = fx.Module("ApiSetting",
	fx.Provide(NewApiConfiguration),
	fx.Provide(func(config ApiConfiguration) Server {
		return *config.Server
	}),
	fx.Provide(func(config ApiConfiguration) Aws {
		return *config.Aws
	}),
	fx.Provide(func(config ApiConfiguration) MongoDb {
		return *config.MongoDb
	}),
)

var WkhtmltopdfAppModule = fx.Module("WkhtmltopdfAppSetting",
	fx.Provide(NewWkhtmltopdfAppConfiguration),
	fx.Provide(func(config WkhtmltopdfAppConfiguration) PdfGeneratorApp {
		return *config.App
	}),
	fx.Provide(func(config WkhtmltopdfAppConfiguration) Aws {
		return *config.Aws
	}),
	fx.Provide(func(config WkhtmltopdfAppConfiguration) S3 {
		return *config.S3
	}),
	fx.Provide(func(config WkhtmltopdfAppConfiguration) SQS {
		return *config.SQS
	}),
	fx.Provide(func(config WkhtmltopdfAppConfiguration) MongoDb {
		return *config.MongoDb
	}),
)

var ChromiumAppModule = fx.Module("ChromiumAppSetting",
	fx.Provide(NewChromiumAppConfiguration),
	fx.Provide(func(config ChromiumAppConfiguration) PdfGeneratorApp {
		return *config.App
	}),
	fx.Provide(func(config ChromiumAppConfiguration) Aws {
		return *config.Aws
	}),
	fx.Provide(func(config ChromiumAppConfiguration) S3 {
		return *config.S3
	}),
	fx.Provide(func(config ChromiumAppConfiguration) SQS {
		return *config.SQS
	}),
	fx.Provide(func(config ChromiumAppConfiguration) MongoDb {
		return *config.MongoDb
	}),
)
