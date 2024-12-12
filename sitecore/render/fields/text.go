package fields

import (
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
	"github.com/mitchellh/mapstructure"
	"guitarrich/xmcloud/model"
)

func RichTextField(fields interface{}, fieldName string) templ.Component {
	field, _ := GetRichTextField(fields, fieldName)
	return RenderRichText(field)
}

func RenderRichText(field model.RichTextField) templ.Component {
	renderedField := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, fmt.Sprintf("%s", field.Value))
		return err
	})

	return RenderFieldMetadata(field.Metadata, renderedField)
}

func GetFieldWithFallback(fields interface{}, fieldNames ...string) string {
	var result model.RichTextField
	var ok bool
	for _, fieldName := range fieldNames {
		result, ok = GetRichTextField(fields, fieldName)
		if ok {
			return fmt.Sprintf("%s", result.Value)
		}
	}

	return ""
}

func GetRichTextField(fields interface{}, fieldName string) (model.RichTextField, bool) {
	fieldMap, ok := fields.(map[string]interface{})
	if !ok {
		fmt.Println("GetRichTextField: not a map")
		return model.RichTextField{}, false
	}
	baseField, ok := fieldMap[fieldName].(map[string]interface{})
	if !ok {
		fmt.Println("GetRichTextField: not a field")
		return model.RichTextField{}, false
	}

	var result model.RichTextField
	err := mapstructure.Decode(baseField, &result)
	if err != nil {
		fmt.Printf("GetRichTextField: not a RichTextField, %s", err)
		return model.RichTextField{}, false
	}

	return result, true
}
