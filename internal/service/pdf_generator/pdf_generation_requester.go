package pdf_generator

import (
	"context"

	"github.com/google/uuid"
	"github.com/mattanapol/kaewsai-pdf/internal/domain"
)

type PdfGenerationRequestServicer interface {
	Request(context context.Context, generator domain.Generator,
		request domain.PdfGenerateRequestForm) error
}

type PdfGenerationRequestService struct {
	pdfGenerateRequester          domain.PdfGenerateRequester
	pdfGenerationRecordRepository domain.PdfGenerationRecordRepository
}

func NewPdfGenerationRequestService(
	pdfGenerateRequester domain.PdfGenerateRequester,
	pdfGenerationRecordRepository domain.PdfGenerationRecordRepository,
) PdfGenerationRequestServicer {
	return &PdfGenerationRequestService{
		pdfGenerateRequester,
		pdfGenerationRecordRepository,
	}
}

func (s *PdfGenerationRequestService) Request(context context.Context,
	generator domain.Generator,
	request domain.PdfGenerateRequestForm) error {
	if request.Id == uuid.Nil {
		request.Id = uuid.New()
	}

	err := s.pdfGenerateRequester.Request(context, generator, request)
	if err != nil {
		return err
	}

	pdfGenerationRecord := domain.NewPdfGenerationRecord(
		request.Id,
		"",
		"",
		"queued",
		string(generator),
	)

	_, err = s.pdfGenerationRecordRepository.Insert(context, pdfGenerationRecord)
	if err != nil {
		return err
	}

	return nil
}
