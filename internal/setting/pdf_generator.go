package setting

// Setting for PDF generator

type PdfGeneratorApp struct {
	OutputPath string `mapstructure:"output-path"`
}

type S3 struct {
	BucketName string `mapstructure:"bucket-name"`
}

type SQS struct {
	InputQueueUrl string `mapstructure:"input-queue-url"`
}

func (s *SQS) GetInputQueueUrl() string {
	return s.InputQueueUrl
}
