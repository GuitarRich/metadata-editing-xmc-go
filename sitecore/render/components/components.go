package components

import (
	"strings"

	"guitarrich/xmcloud/model"
)

func RenderComponentOpen(component model.PlaceholderComponent, context model.SitecoreContext) string {
	var result strings.Builder
	if context.PageEditing {
		result.WriteString("<code type=\"text/sitecore\" chrometype=\"rendering\" class=\"scpm\" kind=\"open\"")
		result.WriteString("id=\"")
		result.WriteString(component.UID)
		result.WriteString("\" style=\"cursor:pointer;\"></code>")
	}

	return result.String()
}

func RenderComponentClose(component model.PlaceholderComponent, context model.SitecoreContext) string {
	if context.PageEditing {
		return "<code type=\"text/sitecore\" chrometype=\"rendering\" class=\"scpm\" kind=\"close\" style=\"cursor:pointer;\"></code>"
	}

	return ""
}
