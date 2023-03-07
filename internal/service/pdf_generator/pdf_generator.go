package pdf_generator

import (
	"context"
	"errors"
	"fmt"

	"github.com/mattanapol/kaewsai-pdf/internal/domain"

	"github.com/google/uuid"
	"github.com/mattanapol/kaewsai-pdf/internal/service/common"
	"github.com/mattanapol/kaewsai-pdf/internal/setting"
)

type PdfGeneratorService struct {
	pdfGeneratorSetting           setting.PdfGeneratorApp
	pdfGenerator                  domain.PdfGenerator
	fileRepository                domain.FileRepository
	pdfGenerationRecordRepository domain.PdfGenerationRecordRepository
}

func NewPdfGeneratorService(
	pdfGeneratorSetting setting.PdfGeneratorApp,
	pdfGenerator domain.PdfGenerator,
	fileRepository domain.FileRepository,
	pdfGenerationRecordRepository domain.PdfGenerationRecordRepository,
) domain.PDfGeneratorServicer {
	return &PdfGeneratorService{pdfGeneratorSetting,
		pdfGenerator,
		fileRepository,
		pdfGenerationRecordRepository}
}

func (s *PdfGeneratorService) Generate(context context.Context, id uuid.UUID, url string, options *domain.PdfGenerateRequestOption) error {
	// Generate pdf
	pdfGenerationRecord := domain.NewPdfGenerationRecord(
		id,
		"",
		"",
		"generating",
		s.pdfGenerator.GeneratorName(),
	)

	_, err := s.pdfGenerationRecordRepository.Update(context, id, pdfGenerationRecord)
	if err != nil {
		return err
	}

	file, err := s.pdfGenerator.GenerateFromLink(context, url, options)
	if err != nil {
		pdfGenerationRecord.Status = fmt.Sprintf("generator failed: %s", err.Error())
		_, repos_err := s.pdfGenerationRecordRepository.Update(context, id, pdfGenerationRecord)
		if repos_err != nil {
			return errors.Join(err, repos_err)
		}

		return err
	}

	// Upload pdf
	fileUploadRequest := domain.FileUploadRequest{
		FileName: uuid.New().String() + common.PdfFileExtension,
		FilePath: s.pdfGeneratorSetting.OutputPath,
		File:     file,
	}

	uploadResponse, err := s.fileRepository.UploadFile(context, fileUploadRequest)

	if err != nil {
		pdfGenerationRecord.Status = fmt.Sprintf("upload failed: %s", err.Error())
		_, repos_err := s.pdfGenerationRecordRepository.Update(context, id, pdfGenerationRecord)
		if repos_err != nil {
			return errors.Join(err, repos_err)
		}

		return err
	}

	pdfGenerationRecord.Status = "success"
	pdfGenerationRecord.FilePath = uploadResponse.FilePath
	pdfGenerationRecord.Bucket = uploadResponse.DriveName
	_, err = s.pdfGenerationRecordRepository.Update(context, id, pdfGenerationRecord)
	return err
}
