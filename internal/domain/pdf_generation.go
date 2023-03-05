package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type PdfGenerationRecord struct {
	ID         uuid.UUID `bson:"id"`
	CreatedOn  time.Time `bson:"createdOn"`
	ModifiedOn time.Time `bson:"modifiedOn"`
	Bucket     string    `bson:"bucket"`
	FilePath   string    `bson:"filePath"`
	Generator  string    `bson:"generator"`
	Status     string    `bson:"status"`
}

func NewPdfGenerationRecord(id uuid.UUID,
	s3Bucket string,
	filePath string,
	status string,
	generator string,
) *PdfGenerationRecord {
	return &PdfGenerationRecord{
		ID:         id,
		CreatedOn:  time.Now(),
		ModifiedOn: time.Now(),
		Bucket:     s3Bucket,
		FilePath:   filePath,
		Status:     status,
		Generator:  generator,
	}
}

func (p *PdfGenerationRecord) UpdateStatus(s3Bucket, filePath, status string) {
	p.Status = status
	p.Bucket = s3Bucket
	p.FilePath = filePath
	p.ModifiedOn = time.Now()
}

type PdfGenerationRecordRepository interface {
	Insert(context context.Context, pdfGenerationRecord *PdfGenerationRecord) (*PdfGenerationRecord, error)
	Update(context context.Context,
		id uuid.UUID,
		pdfGenerationRecord *PdfGenerationRecord) (*PdfGenerationRecord, error)
	FindById(context context.Context, id uuid.UUID) (*PdfGenerationRecord, error)
}

type PdfGenerateRequestForm struct {
	Id      uuid.UUID
	Url     string
	Options PdfGenerateRequestOption
}

type PdfGenerateRequestOption struct {
	Landscape *bool
	Scale     *float32
}

type Generator string

const (
	Wkhtmltopdf Generator = "wkhtmltopdf"
	Chromium    Generator = "chromium"
)

type PdfGenerateRequester interface {
	Request(context context.Context, generator Generator, request PdfGenerateRequestForm) error
}
