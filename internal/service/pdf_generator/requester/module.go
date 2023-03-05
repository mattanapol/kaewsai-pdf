package requester

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/mattanapol/kaewsai-pdf/internal/domain"
	"github.com/mattanapol/kaewsai-pdf/internal/service/aws"
	"github.com/mattanapol/kaewsai-pdf/internal/setting"
	"go.uber.org/fx"
)

var Module = fx.Module("PdfGeneratorRequestService",
	fx.Provide(aws.NewSqsClient),
	fx.Provide(func(sqsClient *sqs.Client,
		apiConfiguration setting.ApiConfiguration,
	) domain.PdfGenerateRequester {
		return NewPdfGenerateRequestService(sqsClient, apiConfiguration.Wk, apiConfiguration.Chromium)
	}),
)
