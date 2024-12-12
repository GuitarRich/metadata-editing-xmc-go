package editing

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
	"guitarrich/xmcloud/sitecore/render/components"
)

func (h *EditingRequestHandler) Config(c echo.Context) error {

	fmt.Println("editingRenderMiddlewareConfig")
	fmt.Printf("Request: %s\n", c.Request().URL.Path)

	// Validate allowed origins and editing secret
	if !enforceCors(c, EDITING_ALLOWED_ORIGINS) {
		message := fmt.Sprintf("Requests from origin %s not allowed", c.Request().Header.Get("Origin"))
		return c.HTML(http.StatusForbidden, message)
	}

	// Validate editing secret
	secret := c.Request().URL.Query().Get(QUERY_PARAM_EDITING_SECRET)
	editingSecret, err := getJssEditingSecret()
	if err != nil {
		return c.HTML(http.StatusForbidden, err.Error())
	}

	if secret != editingSecret {
		message := fmt.Sprintf("Invalid editing secret")
		return c.HTML(http.StatusForbidden, message)
	}

	c.Response().Header().Set("Content-Security-Policy", getSCHeader())

	packages := make(map[string]interface{})

	componentList := components.GetComponents()
	keys := make([]string, 0, len(componentList))

	for key := range componentList {
		keys = append(keys, key)
	}

	sort.Sort(sort.StringSlice(keys))

	response := EditingMiddlewareConfig{
		Components:    keys,
		PagesEditMode: EditModeMetadata,
		Packages:      packages,
	}

	fmt.Println("editingRenderMiddlewareConfig - END")
	fmt.Printf("Response: %s\n", c.Response().Header().Get("Content-Type"))
	return c.JSON(http.StatusOK, response)
}
