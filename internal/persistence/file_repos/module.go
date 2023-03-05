package file_repos

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mattanapol/kaewsai-pdf/internal/service/aws"
	"github.com/mattanapol/kaewsai-pdf/internal/setting"
	"go.uber.org/fx"
)

var S3FileModule = fx.Module("S3FileRepository",
	fx.Provide(aws.NewS3Client),
	fx.Provide(func(s3Client *s3.Client, s3Setting setting.S3) FileRepository {
		return NewS3Repository(s3Client, s3Setting)
	},
	),
)

type FileRepository interface {
	UploadFile(context context.Context, fileUploadRequest FileUploadRequest) (*FileUploadResponse, error)
}

type FileUploadRequest struct {
	FileName string
	FilePath string
	File     io.Reader
	Metadata map[string]string
}

type FileUploadResponse struct {
	DriveName string
	FilePath  string
}
