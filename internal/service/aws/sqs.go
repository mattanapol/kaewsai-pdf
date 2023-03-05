package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/mattanapol/kaewsai-pdf/internal/setting"
)

func NewSqsClient(awsSetting setting.Aws) *sqs.Client {
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

	return sqs.NewFromConfig(conf)
}
