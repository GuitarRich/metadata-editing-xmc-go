package fields

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"guitarrich/xmcloud/model"
)

func renderFieldMetadata(metadata model.MetadataData, field templ.Component) templ.Component {
	fieldMetadata := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return RenderFieldMetadata(metadata, field).Render(ctx, w)
	})
	return fieldMetadata
}
