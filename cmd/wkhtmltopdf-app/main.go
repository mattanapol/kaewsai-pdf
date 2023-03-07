package main

import (
	"context"
	"log"

	"github.com/mattanapol/kaewsai-pdf/internal/domain"
	"github.com/mattanapol/kaewsai-pdf/internal/persistence/file_repos"
	"github.com/mattanapol/kaewsai-pdf/internal/persistence/nosql_repos"

	"github.com/mattanapol/kaewsai-pdf/internal/service/pdf_generator"
	"github.com/mattanapol/kaewsai-pdf/internal/service/pdf_generator/receiver"
	"github.com/mattanapol/kaewsai-pdf/internal/service/pdf_generator/wkhtmltopdf_generator"
	"github.com/mattanapol/kaewsai-pdf/internal/setting"
	"go.uber.org/fx"
)

func main() {
	// Run server
	app := fx.New(
		wkhtmltopdf_generator.Module,
		setting.WkhtmltopdfAppModule,
		nosql_repos.Module,
		pdf_generator.Module,
		receiver.Module,
		file_repos.S3FileModule,
		fx.Invoke(
			startApplication,
		),
		// fx.NopLogger, // To hide dependency injection log
	)

	app.Run()
}

func startApplication(lc fx.Lifecycle, pdfGeneratorReceiver domain.PdfGenerateRequestReceiver) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Start Application ...")
			// Create context from parent context
			go func() {
				// While loop
				for {
					log.Println("Waiting for message ...")
					err := pdfGeneratorReceiver.Receive(context.Background())
					if err != nil {
						log.Printf("Error: %s", err)
					}
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutdown Application ...")

			return nil
		},
	})
}
