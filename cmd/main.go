package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"guitarrich/xmcloud/middleware/editing"
	"guitarrich/xmcloud/sitecore"
	"guitarrich/xmcloud/sitecore/handlers"
	"guitarrich/xmcloud/sitecore/pipelines/requestBegin"
	"guitarrich/xmcloud/views/components"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"
)

func main() {
	app := echo.New()

	app.Static("/favicon.ico", "favicon.ico")
	app.Static("/css", "css")
	app.Static("/js", "js")
	app.Static("/img", "img")

	app.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	app.Use(middleware.Logger())

	components.RegisterComponents()

	// Kick off the main request pipeline
	app.GET("/*", requestBegin.NewRequestPipelineHandler().RunPipeline)

	// Register the editng API endpoints
	app.OPTIONS("/api/editing/config", editing.NewEditingRequestHandler().Config)
	app.GET("/api/editing/config", editing.NewEditingRequestHandler().Config)
	app.GET("/api/editing/render", editing.NewEditingRequestHandler().Render)

	// handle the XM `/-/media` endpoint
	app.GET("/-/media", handlers.NewMediaHandler().ServeHTTP)

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	port := sitecore.GetEnvVar("PORT")
	if err := app.StartH2CServer(fmt.Sprintf(":%s", port), s); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
