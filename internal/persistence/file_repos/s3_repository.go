package file_repos

import (
	"context"
	"log"
	"path/filepath"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mattanapol/kaewsai-pdf/internal/setting"
)

type S3Repository struct {
	client    *s3.Client
	s3Setting setting.S3
}

func NewS3Repository(s3Client *s3.Client, s3Setting setting.S3) *S3Repository {
	log.Printf("S3 BucketName: %s", s3Setting.BucketName)
	return &S3Repository{
		client:    s3Client,
		s3Setting: s3Setting,
	}
}

func (r *S3Repository) UploadFile(context context.Context,
	fileUploadRequest FileUploadRequest,
) (*FileUploadResponse, error) {
	filePath := filepath.Join(fileUploadRequest.FilePath, fileUploadRequest.FileName)
	log.Printf("Uploading file to S3 bucket: %s, File: %s",
		r.s3Setting.BucketName,
		filePath,
	)
	params := &s3.PutObjectInput{
		Bucket:   aws.String(r.s3Setting.BucketName),
		Key:      aws.String(filePath),
		Body:     fileUploadRequest.File,
		Metadata: generateMetadata(fileUploadRequest.Metadata),
	}

	_, err := r.client.PutObject(context, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &FileUploadResponse{
		DriveName: r.s3Setting.BucketName,
		FilePath:  filePath,
	}, nil
}

func generateMetadata(inputMetadata map[string]string) map[string]string {
	metadata := make(map[string]string)
	for key, value := range inputMetadata {
		metadata[key] = value
	}
	if _, ok := metadata["uploadTimestamp"]; !ok {
		metadata["uploadTimestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	}
	return metadata
}
