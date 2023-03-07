package file_repos

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mattanapol/kaewsai-pdf/internal/domain"
	"github.com/mattanapol/kaewsai-pdf/internal/service/aws"
	"github.com/mattanapol/kaewsai-pdf/internal/setting"
	"go.uber.org/fx"
)

var S3FileModule = fx.Module("S3FileRepository",
	fx.Provide(aws.NewS3Client),
	fx.Provide(func(s3Client *s3.Client, s3Setting setting.S3) domain.FileRepository {
		return NewS3Repository(s3Client, s3Setting)
	},
	),
)
