package domain

import (
	"context"
	"io"

	"github.com/google/uuid"
)

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

type PdfGenerator interface {
	GenerateFromLink(context context.Context, url string, Options *PdfGenerateRequestOption) (io.Reader, error)
	GeneratorName() string
}

type PdfGenerateRequester interface {
	Request(context context.Context, generator Generator, request PdfGenerateRequestForm) error
}

type PdfGenerateRequestReceiver interface {
	Receive(context context.Context) error
}

type PDfGeneratorServicer interface {
	Generate(context context.Context, id uuid.UUID, url string, options *PdfGenerateRequestOption) error
}
