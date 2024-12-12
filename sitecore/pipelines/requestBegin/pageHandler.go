package requestBegin

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"guitarrich/xmcloud/model"
	"guitarrich/xmcloud/sitecore"
)

func (h *RequestPipelineHandler) handlePage(c echo.Context) error {

	sitecore.RendererLog("RequestPipelineHandler\nRequest: %s\n", c.Request().URL.Path)

	// are we in preview mode?
	var query string
	var queryName string
	headers := map[string]string{}
	params := map[string]string{}

	sitecore.RendererLog("Query: %s\n", c.Request().URL.RawQuery)
	sitecore.RendererLog("mode: %s\n", c.Request().URL.Query().Get("mode"))

	sitecore.RendererLog("Normal mode\n")
	queryName = "LayoutQuery"
	query = sitecore.GetLayoutQuery(h.itemPath, h.language, h.siteName)

	result := sitecore.RunQueryWithParameters(query, queryName, headers, params)

	jsonString, _ := json.Marshal(result)

	layoutResponse := model.LayoutResponse{}
	json.Unmarshal(jsonString, &layoutResponse)

	if layoutResponse.Data.Layout.Item.Rendered.Sitecore.Route.Placeholders == nil {
		// This is a 404, so we need to render the 404 page
		sitecore.RendererLog("404: NotFound\n")
		h.NotFound = true
		return nil
	}

	h.renderedLayout = layoutResponse.Data.Layout.Item.Rendered

	return nil
}
