package fields

import (
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
	"github.com/mitchellh/mapstructure"
	"guitarrich/xmcloud/model"
	"guitarrich/xmcloud/sitecore"
)

func LinkField(fields interface{}, fieldName string, classNames ...string) templ.Component {
	field := GetLinkField(fields, fieldName)
	return renderFieldMetadata(field.Metadata, renderLink(field, classNames...))
}

func LinkFieldHasLink(fields interface{}, fieldName string) bool {
	field := GetLinkField(fields, fieldName)
	return field.Value.Href != ""
}

func renderLink(field model.LinkField, classNames ...string) templ.Component {
	href := field.Value.Href
	if field.Value.Querystring != "" {
		href += "?" + field.Value.Querystring
	}
	if field.Value.Anchor != "" {
		href += "#" + field.Value.Anchor
	}

	link := fmt.Sprintf("<a  href=\"%s\"", href)
	link += sitecore.AddIfNotEmpty("target", field.Value.Target)
	link += sitecore.AddIfNotEmpty("title", field.Value.Title)

	cssClasses := field.Value.Class
	for _, className := range classNames {
		cssClasses += " " + className
	}
	link += sitecore.AddIfNotEmpty("class", cssClasses)

	if field.Value.Target == "_blank" {
		link += " rel=\"noopener noreferrer\""
	}

	link += fmt.Sprintf(">%s</a>", field.Value.Text)

	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, link)
		return err
	})
}

func GetLinkField(fields interface{}, fieldName string) model.LinkField {
	fieldMap, ok := fields.(map[string]interface{})
	if !ok {
		fmt.Println("GetLinkField: not a map")
		return model.LinkField{}
	}

	baseField, ok := fieldMap[fieldName].(map[string]interface{})
	if !ok {
		fmt.Println("GetLinkField: not a field")
		return model.LinkField{}
	}

	var result model.LinkField

	err := mapstructure.Decode(baseField, &result)
	if err != nil {
		fmt.Printf("GetLinkField: not a LinkField, %s", err)
		return model.LinkField{}
	}

	return result
}
