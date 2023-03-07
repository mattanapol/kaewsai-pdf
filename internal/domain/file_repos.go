package domain

import (
	"context"
	"io"
)

type FileRepository interface {
	UploadFile(context context.Context, fileUploadRequest FileUploadRequest) (*FileUploadResponse, error)
}

type FileUploadRequest struct {
	FileName string
	FilePath string
	File     io.Reader
	Metadata map[string]string
}

type FileUploadResponse struct {
	DriveName string
	FilePath  string
}
