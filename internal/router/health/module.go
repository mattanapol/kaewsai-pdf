package health

import "go.uber.org/fx"

var Module = fx.Module("HealthControllers",
	fx.Provide(NewHealthController),
)
