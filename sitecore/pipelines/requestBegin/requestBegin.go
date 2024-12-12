package requestBegin

import (
	"net/http"

	"guitarrich/xmcloud/model"
	"guitarrich/xmcloud/sitecore"
	"guitarrich/xmcloud/views/layout"

	"github.com/labstack/echo/v4"
)

type RequestPipelineHandler struct {
	siteName       string
	language       string
	itemPath       string
	itemId         string
	context        echo.Context
	renderedLayout model.Rendered
	NotFound       bool
}

func NewRequestPipelineHandler() *RequestPipelineHandler {
	return &RequestPipelineHandler{}
}

func (h *RequestPipelineHandler) RunPipeline(c echo.Context) error {

	h.siteName = sitecore.GetEnvVar("SITECORE_SITE_NAME")
	h.language = sitecore.GetEnvVar("SITECORE_LANGUAGE")
	h.itemPath = c.Request().URL.Path

	// Handle special cases
	if h.itemPath == "/sitemap.xml" {
		return h.handleSitemap(c)
	}

	// Handle catchall paths
	_ = h.handlePage(c)
	if h.NotFound {
		// Check for redirects
		_ = h.handleRedirects(c)
		if h.NotFound {
			_ = h.handleNotFound(c)
		}
	}

	var tmp model.PlaceholderComponent
	handleDynamicPlaceholders(tmp, h.renderedLayout.Sitecore.Route.Placeholders, 1)

	response := render(c, http.StatusOK, layout.MainLayout(h.renderedLayout))
	return response
}
