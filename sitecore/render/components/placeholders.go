package components

import (
	"fmt"
	"strings"

	"guitarrich/xmcloud/model"
)

func RenderPlaceholderOpen(placeholderKey string, id string, context model.SitecoreContext) string {
	if id == "" {
		id = "00000000-0000-0000-0000-000000000000"
	}
	var result strings.Builder
	if context.PageEditing {
		result.WriteString("<code type=\"text/sitecore\" chrometype=\"placeholder\" class=\"scpm\" kind=\"open\"")
		result.WriteString("id=\"")
		result.WriteString(fmt.Sprintf("%s_%s", placeholderKey, id))
		result.WriteString("\" style=\"cursor:pointer;\"></code>")
	}

	return result.String()
}

func RenderPlaceholderClose(placeholderKey string, context model.SitecoreContext) string {
	if context.PageEditing {
		return "<code type=\"text/sitecore\" chrometype=\"placeholder\" class=\"scpm\" kind=\"close\" style=\"cursor:pointer;\"></code>"
	}

	return ""
}
