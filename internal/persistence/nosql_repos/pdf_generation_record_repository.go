package nosql_repos

import (
	"context"

	"github.com/google/uuid"
	"github.com/mattanapol/kaewsai-pdf/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "pdf_generation_record"

type pdfGenerationRecordRepositoryImpl struct {
	client *mongo.Collection
}

func NewPdfGenerationRecordRepository(mongoDb *mongo.Database) domain.PdfGenerationRecordRepository {
	return &pdfGenerationRecordRepositoryImpl{
		client: mongoDb.Collection(collectionName),
	}
}

func (r *pdfGenerationRecordRepositoryImpl) Insert(context context.Context,
	pdfGenerationRecord *domain.PdfGenerationRecord) (*domain.PdfGenerationRecord, error) {
	if pdfGenerationRecord.ID == uuid.Nil {
		pdfGenerationRecord.ID = uuid.New()
	}
	_, err := r.client.InsertOne(context, pdfGenerationRecord)
	if err != nil {
		return nil, err
	}
	return pdfGenerationRecord, nil
}

func (r *pdfGenerationRecordRepositoryImpl) Update(context context.Context,
	id uuid.UUID,
	pdfGenerationRecord *domain.PdfGenerationRecord) (*domain.PdfGenerationRecord, error) {
	if id == uuid.Nil {
		return nil, &domain.MissingIdError{}
	}

	var (
		filter = bson.D{{Key: "id", Value: id}}
	)
	pdfGenerationRecordDb, err := r.FindById(context, id)
	if err != nil {
		return nil, err
	}

	pdfGenerationRecordDb.ModifiedOn = pdfGenerationRecord.ModifiedOn
	pdfGenerationRecordDb.Status = pdfGenerationRecord.Status
	pdfGenerationRecordDb.Bucket = pdfGenerationRecord.Bucket
	pdfGenerationRecordDb.FilePath = pdfGenerationRecord.FilePath
	pdfGenerationRecordDb.Generator = pdfGenerationRecord.Generator

	_, err = r.client.UpdateOne(
		context,
		filter,
		bson.D{{Key: "$set", Value: pdfGenerationRecordDb}},
	)
	if err != nil {
		return nil, err
	}

	return pdfGenerationRecordDb, nil
}

func (r *pdfGenerationRecordRepositoryImpl) FindById(context context.Context, id uuid.UUID) (*domain.PdfGenerationRecord, error) {
	if id == uuid.Nil {
		return nil, &domain.MissingIdError{}
	}

	var (
		pdfGenerationRecordDb *domain.PdfGenerationRecord
		filter                = bson.D{{Key: "id", Value: id}}
	)
	findResult := r.client.FindOne(context, filter)
	if findResult.Err() != nil {
		return nil, findResult.Err()
	}

	err := findResult.Decode(&pdfGenerationRecordDb)
	if err != nil {
		return nil, err
	}

	return pdfGenerationRecordDb, nil
}
