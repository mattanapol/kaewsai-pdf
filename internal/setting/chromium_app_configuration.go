package setting

type ChromiumAppConfiguration struct {
	App     *PdfGeneratorApp `mapstructure:"app"`
	MongoDb *MongoDb         `mapstructure:"mongodb"`
	Aws     *Aws             `mapstructure:"aws"`
	S3      *S3              `mapstructure:"s3"`
	SQS     *SQS             `mapstructure:"sqs"`
}

// NewChromiumAppConfiguration initialize the configuration instance
func NewChromiumAppConfiguration() (ChromiumAppConfiguration, error) {
	cfg := ChromiumAppConfiguration{
		App:     &PdfGeneratorApp{},
		MongoDb: &MongoDb{},
		Aws:     &Aws{},
		S3:      &S3{},
		SQS:     &SQS{},
	}

	readFile(cfg, "config", "chromium-app", "")

	return cfg, nil
}
