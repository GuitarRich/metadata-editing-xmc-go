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

func ImageField(fields interface{}, fieldName string, cssClass string) templ.Component {
	field := GetImageField(fields, fieldName)
	return renderFieldMetadata(field.Metadata, renderImage(field, cssClass))
}

func renderImage(field model.ImageField, cssClass string) templ.Component {
	img := fmt.Sprintf("<img src=\"%s\"", field.Value.Src)
	img += sitecore.AddIfNotEmpty("alt", field.Value.Alt)
	img += sitecore.AddIfNotEmpty("width", field.Value.Width)
	img += sitecore.AddIfNotEmpty("height", field.Value.Height)

	if cssClass != "" {
		img += fmt.Sprintf(" class=\"%s\"", cssClass)
	}

	img += fmt.Sprintf(" />")

	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, img)
		return err
	})
}

func GetImageField(fields interface{}, fieldName string) model.ImageField {
	fieldMap, ok := fields.(map[string]interface{})
	if !ok {
		fmt.Println("GetImageField: not a map")
		return model.ImageField{}
	}
	baseField, ok := fieldMap[fieldName].(map[string]interface{})
	if !ok {
		fmt.Println("GetImageField: not a field")
		return model.ImageField{}
	}

	var imageField model.ImageField
	err := mapstructure.Decode(baseField, &imageField)
	if err != nil {
		fmt.Printf("GetImageField: not an ImageField, %s", err)
		return model.ImageField{}
	}

	return imageField
}
