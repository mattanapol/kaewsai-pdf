package wkhtmltopdf_generator

import (
	"bytes"
	"context"
	"io"
	"log"
	"time"

	"github.com/mattanapol/kaewsai-pdf/internal/domain"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type WkhtmltopdfGenerator struct {
	fileRepository domain.FileRepository
}

func NewWkhtmltopdfGenerator(fileRepository domain.FileRepository) WkhtmltopdfGenerator {
	return WkhtmltopdfGenerator{
		fileRepository: fileRepository,
	}
}

func (wk WkhtmltopdfGenerator) GeneratorName() string {
	return string(domain.Wkhtmltopdf)
}

func (wk WkhtmltopdfGenerator) GenerateFromLink(context context.Context,
	url string,
	options *domain.PdfGenerateRequestOption) (io.Reader, error) {
	// Create new PDF generator
	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Set global options
	setPdfGeneratorFromOptions(pdfGenerator, options)

	// Create a new input page from an URL
	page := wkhtmltopdf.NewPage(url)

	// Set options for this page
	setPageGeneratorFromOptions(page, options)

	// Add to document
	pdfGenerator.AddPage(page)

	errBuf := new(bytes.Buffer)
	pdfGenerator.SetStderr(errBuf)

	done := false
	defer func() { done = true }()
	go func() {
		for !done {
			time.Sleep(500 * time.Millisecond)
			log.Println(errBuf.String())
		}
	}()

	buffer := &bytes.Buffer{}
	pdfGenerator.SetOutput(buffer)
	// Create PDF document in internal buffer
	log.Println("Creating PDF file...")
	err = pdfGenerator.Create()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("PDF file generated successfully")

	return buffer, nil
}

func setPdfGeneratorFromOptions(pdfGenerator *wkhtmltopdf.PDFGenerator, options *domain.PdfGenerateRequestOption) {
	pdfGenerator.Dpi.Set(300)
	pdfGenerator.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	if options == nil {
		return
	}
	if options.Landscape != nil {
		if *options.Landscape {
			pdfGenerator.Orientation.Set(wkhtmltopdf.OrientationLandscape)
		} else {
			pdfGenerator.Orientation.Set(wkhtmltopdf.OrientationPortrait)
		}
	}
}

func setPageGeneratorFromOptions(page *wkhtmltopdf.Page, options *domain.PdfGenerateRequestOption) {
	page.NoStopSlowScripts.Set(true)
	page.DisableJavascript.Set(true)

	if options == nil {
		return
	}
	if options.Scale != nil {
		page.Zoom.Set(float64(*options.Scale))
	}
}
