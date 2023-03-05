package requester

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/mattanapol/kaewsai-pdf/internal/domain"
)

type PdfGenerateRequestService struct {
	sqsClient                *sqs.Client
	wkhtmltopdfRequestGetter WkhtmltopdfRequestGetter
	chromiumRequestGetter    ChromiumRequestGetter
}

type RequestGetter interface {
	GetInputQueueUrl() string
}
type WkhtmltopdfRequestGetter interface {
	RequestGetter
}
type ChromiumRequestGetter interface {
	RequestGetter
}

func NewPdfGenerateRequestService(sqsClient *sqs.Client,
	wkhtmltopdfRequestGetter WkhtmltopdfRequestGetter,
	chromiumRequestGetter ChromiumRequestGetter,
) *PdfGenerateRequestService {
	return &PdfGenerateRequestService{
		sqsClient:                sqsClient,
		wkhtmltopdfRequestGetter: wkhtmltopdfRequestGetter,
		chromiumRequestGetter:    chromiumRequestGetter,
	}
}

func (service *PdfGenerateRequestService) Request(context context.Context,
	generator domain.Generator,
	request domain.PdfGenerateRequestForm) error {
	// marshal request to json string
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}
	jsonString := string(jsonBytes)

	var queueUrl string
	if generator == domain.Wkhtmltopdf {
		queueUrl = service.wkhtmltopdfRequestGetter.GetInputQueueUrl()
	} else if generator == domain.Chromium {
		queueUrl = service.chromiumRequestGetter.GetInputQueueUrl()
	} else {
		return &domain.InvalidRequestTypeError{}
	}

	// send to sqs
	output, err := service.sqsClient.SendMessage(context, &sqs.SendMessageInput{
		MessageBody: &jsonString,
		QueueUrl:    &queueUrl,
	})
	log.Println(output)

	return err
}
