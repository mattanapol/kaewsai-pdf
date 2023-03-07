package receiver

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/mattanapol/kaewsai-pdf/internal/domain"
	"github.com/mattanapol/kaewsai-pdf/internal/service/aws"
	"github.com/mattanapol/kaewsai-pdf/internal/setting"
	"go.uber.org/fx"
)

var Module = fx.Module("PdfGeneratorReceiver",
	fx.Provide(aws.NewSqsClient),
	fx.Provide(func(sqsClient *sqs.Client,
		sqsSetting setting.SQS,
		pdfGeneratorService domain.PDfGeneratorServicer) domain.PdfGenerateRequestReceiver {
		return NewPdfGenerateRequestReceiver(sqsClient, sqsSetting, pdfGeneratorService)
	}),
)
