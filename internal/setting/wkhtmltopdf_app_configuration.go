package setting

type WkhtmltopdfAppConfiguration struct {
	App     *PdfGeneratorApp `mapstructure:"app"`
	MongoDb *MongoDb         `mapstructure:"mongodb"`
	Aws     *Aws             `mapstructure:"aws"`
	S3      *S3              `mapstructure:"s3"`
	SQS     *SQS             `mapstructure:"sqs"`
}

// NewWkhtmltopdfAppConfiguration initialize the configuration instance
func NewWkhtmltopdfAppConfiguration() (WkhtmltopdfAppConfiguration, error) {
	cfg := WkhtmltopdfAppConfiguration{
		App:     &PdfGeneratorApp{},
		MongoDb: &MongoDb{},
		Aws:     &Aws{},
		S3:      &S3{},
		SQS:     &SQS{},
	}

	readFile(cfg, "config", "wkhtmltopdf-app", "")

	return cfg, nil
}
