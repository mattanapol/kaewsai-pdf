package receiver

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/mattanapol/kaewsai-pdf/internal/domain"
	"github.com/mattanapol/kaewsai-pdf/internal/setting"
)

type PdfGenerateRequestReceiverService struct {
	sqsClient           *sqs.Client
	sqsSetting          setting.SQS
	pdfGeneratorService domain.PDfGeneratorServicer
}

func NewPdfGenerateRequestReceiver(sqsClient *sqs.Client,
	sqsSetting setting.SQS,
	pdfGenerator domain.PDfGeneratorServicer) *PdfGenerateRequestReceiverService {
	if sqsSetting.InputQueueUrl == "" {
		log.Panicln("SQS input queue url is not set")
	}
	return &PdfGenerateRequestReceiverService{
		sqsClient:           sqsClient,
		pdfGeneratorService: pdfGenerator,
		sqsSetting:          sqsSetting,
	}
}

func (service *PdfGenerateRequestReceiverService) Receive(context context.Context) error {
	queueUrl := service.sqsSetting.InputQueueUrl

	messages, err := service.sqsClient.ReceiveMessage(context, &sqs.ReceiveMessageInput{
		QueueUrl:            &queueUrl,
		MaxNumberOfMessages: 3,
		WaitTimeSeconds:     20,
	})

	if err != nil {
		return err
	}

	for _, message := range messages.Messages {
		var request domain.PdfGenerateRequestForm
		err := json.Unmarshal([]byte(*message.Body), &request)
		if err != nil {
			return err
		}
		log.Printf("Received message: %s, %v", *message.MessageId, request)
		err = service.pdfGeneratorService.Generate(context, request.Id, request.Url, &request.Options)
		if err != nil {
			log.Println(err)
		} else {
			// delete message
			_, err := service.sqsClient.DeleteMessage(context, &sqs.DeleteMessageInput{
				QueueUrl:      &queueUrl,
				ReceiptHandle: message.ReceiptHandle,
			})
			if err != nil {
				return err
			}
			log.Printf("Done: %s", *message.MessageId)
		}
	}

	return nil
}
