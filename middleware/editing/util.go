package editing

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/labstack/echo/v4"
	"guitarrich/xmcloud/model"
	"guitarrich/xmcloud/sitecore"
)

func NewEditingRequestHandler() *EditingRequestHandler {
	return &EditingRequestHandler{}
}

func enforceCors(c echo.Context, allowedOrigins []string) bool {

	// origin is not present for non-CORS requests (e.g. server-side) - so we skip the checks
	origin := c.Request().Header.Get("Origin")
	if origin == "" {
		fmt.Printf("Origin header not present\n")
		origin = "https://pages.sitecorecloud.io"
		//return true
	}

	// 3 sources of allowed origins are considered:
	// the env value
	defaultAllowedOrigins := getAllowedOriginsFromEnv()
	// the allowedOriigns props
	allowedOrigins = append(allowedOrigins, defaultAllowedOrigins...)

	// and the existing CORS header, if present (i.e set by go config)
	presetCors := c.Request().Header.Get("Access-Control-Allow-Origin")
	if presetCors != "" {
		allowedOrigins = append(allowedOrigins, presetCors)
	}

	idx := slices.IndexFunc(allowedOrigins, func(s string) bool {
		result := s == origin
		if !result {
			r := regexp.MustCompile(s)
			result = r.MatchString(origin)
		}
		return result
	})

	if idx == -1 {
		fmt.Printf("Origin %s is not allowed\n", origin)
		return false
	}

	fmt.Printf("Allowed origins: %s\n", strings.Join(allowedOrigins, ", "))
	c.Response().Header().Set("Access-Control-Allow-Origin", origin)
	c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

	return true
}

func getAllowedOriginsFromEnv() []string {
	allowedOrigins := sitecore.GetEnvVar("JSS_ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		return []string{}
	}
	return strings.Split(allowedOrigins, ",")
}

func convertToWildcardReged(pattern string) string {
	replacedPattern := regexp.MustCompile(`/`).ReplaceAllString(pattern, `\/`)
	replacedPattern = regexp.MustCompile(`\.`).ReplaceAllString(replacedPattern, `\.`)
	replacedPattern = regexp.MustCompile(`\*`).ReplaceAllString(replacedPattern, `.*`)
	return "^" + replacedPattern + "$"
}

func getJssEditingSecret() (string, error) {
	secret := sitecore.GetEnvVar("JSS_EDITING_SECRET")
	if secret == "" {
		return "", errors.New("JSS_EDITING_SECRET is not set")
	}
	return secret, nil
}

func getSCHeader() string {
	var envOrigins = strings.Join(getAllowedOriginsFromEnv(), " ")
	var editingOrigins = strings.Join(EDITING_ALLOWED_ORIGINS, ", ")
	fmt.Printf("envOrigins: %s\n", envOrigins)
	fmt.Printf("editingOrigins: %s\n", editingOrigins)
	return fmt.Sprintf("frame-ancestors 'self' %s %s", envOrigins, editingOrigins)
}

func GetEditScripts(data model.Rendered) string {

	if !sitecore.IsEditMode(data.Sitecore.Context) {
		return ""
	}

	var scripts strings.Builder

	for _, script := range data.Sitecore.Context.ClientScripts {
		scripts.WriteString(fmt.Sprintf("<script src=\"%s\"></script>", script))
	}

	if data.Sitecore.Context.ClientData.HorizonCanvasState.ItemId != "" {
		jsonString, err_ := json.Marshal(data.Sitecore.Context.ClientData.HorizonCanvasState)
		if err_ == nil {
			scripts.WriteString("<script id=\"hrz-canvas-state\" type=\"application/json\">")
			scripts.WriteString(string(jsonString))
			scripts.WriteString("</script>")
		}
	}

	return scripts.String()
}
