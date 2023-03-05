package chromium_generator

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/mattanapol/kaewsai-pdf/internal/domain"
	"github.com/mattanapol/kaewsai-pdf/internal/persistence/file_repos"

	"github.com/go-resty/resty/v2"
)

const chromiumUrl = "http://localhost:3000/forms/chromium/convert/url"

type ChromiumGenerator struct {
	fileRepository file_repos.FileRepository
}

func NewChromiumGenerator(fileRepository file_repos.FileRepository) ChromiumGenerator {
	return ChromiumGenerator{
		fileRepository: fileRepository,
	}
}

func (ch ChromiumGenerator) GeneratorName() string {
	return string(domain.Chromium)
}

func (ch ChromiumGenerator) GenerateFromLink(context context.Context,
	url string,
	options *domain.PdfGenerateRequestOption) (io.Reader, error) {
	// Create a new resty client
	client := resty.New()

	// Define the payload to send in the request
	payload := createPayloadFromOptions(options)
	payload["url"] = url

	// Make the POST request
	response, err := client.R().
		SetDoNotParseResponse(true).
		SetMultipartFormData(payload).
		Post(chromiumUrl)

	// Check if there was an error making the request
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}

	// Print the response status code
	log.Printf("Response Status Code: %d\n", response.StatusCode())

	// Read the response body into a bytes.Buffer
	buffer := &bytes.Buffer{}
	_, err = buffer.ReadFrom(response.RawBody())
	defer response.RawBody().Close()
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return buffer, nil
}

func createPayloadFromOptions(options *domain.PdfGenerateRequestOption) map[string]string {
	payload := createDefaultPayload()
	if options == nil {
		return payload
	}
	if options.Scale != nil {
		payload["scale"] = fmt.Sprint(*options.Scale)
		log.Printf("Scale: %s\n", payload["scale"])
	}
	if options.Landscape != nil {
		payload["landscape"] = fmt.Sprint(*options.Landscape)
		log.Printf("landscape: %s\n", payload["landscape"])
	}
	return payload
}

func createDefaultPayload() map[string]string {
	return map[string]string{
		"marginTop":    "0",
		"marginBottom": "0",
		"marginLeft":   "0",
		"marginRight":  "0",
		"landscape":    "true",
		"scale":        "1.0",
	}
}
