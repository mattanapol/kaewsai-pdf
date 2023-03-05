package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mattanapol/kaewsai-pdf/internal/persistence/nosql_repos"
	"github.com/mattanapol/kaewsai-pdf/internal/router"
	"github.com/mattanapol/kaewsai-pdf/internal/router/health"
	"github.com/mattanapol/kaewsai-pdf/internal/router/pdf_generator"
	"github.com/mattanapol/kaewsai-pdf/internal/setting"
	"go.uber.org/fx"
)

// @title Pdf Generator Service API
// @version 1.0
// @description Pdf Generator Service API
func main() {
	// Run server
	app := fx.New(
		health.Module,
		setting.ApiModule,
		nosql_repos.Module,
		pdf_generator.Module,
		fx.Provide(
			NewServer,
		),
		fx.Invoke(router.InitRouter),
		// fx.NopLogger, // To hide dependency injection log
	)

	app.Run()
}

func NewServer(lc fx.Lifecycle, serverSetting setting.Server) *http.Server {
	gin.SetMode(serverSetting.RunMode)

	readTimeout := serverSetting.ReadTimeout
	writeTimeout := serverSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", serverSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				// service connections
				listener, err := net.Listen("tcp", server.Addr)
				if err != nil {
					log.Printf("listen: %s\n", err)
				}

				if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
					log.Printf("Server serve error: %s\n", err)
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			log.Println("Shutdown Server ...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}
			log.Println("Server exiting")
			return nil
		},
	})

	return server
}
