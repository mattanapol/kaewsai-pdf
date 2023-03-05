package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mattanapol/kaewsai-pdf/internal/setting"
)

func NewS3Client(awsSetting setting.Aws) *s3.Client {
	conf, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(awsSetting.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				awsSetting.AccessKeyId,
				awsSetting.SecretAccessKey,
				"",
			),
		),
	)
	if err != nil {
		panic(err)
	}

	s3Client := s3.NewFromConfig(conf)
	return s3Client
}
