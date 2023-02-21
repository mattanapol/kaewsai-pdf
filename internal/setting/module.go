package setting

import "go.uber.org/fx"

var ApiModule = fx.Module("ApiSetting",
	fx.Provide(NewApiConfiguration),
	fx.Provide(func(config ApiConfiguration) Server {
		return *config.Server
	}),
)
