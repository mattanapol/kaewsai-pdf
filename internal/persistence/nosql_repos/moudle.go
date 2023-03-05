package nosql_repos

import "go.uber.org/fx"

var Module = fx.Module("Persistence",
	fx.Provide(SetupMongoDb),
	fx.Provide(NewPdfGenerationRecordRepository),
)
