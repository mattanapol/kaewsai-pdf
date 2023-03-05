package setting

// Setting for AWS

type Aws struct {
	Region          string `mapstructure:"region"`
	AccessKeyId     string `mapstructure:"access-key-id"`
	SecretAccessKey string `mapstructure:"secret-access-key"`
}
