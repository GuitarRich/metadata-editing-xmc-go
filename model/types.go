package model

type EditMode string

const (
	Chromes  EditMode = "chromes"
	Metadata          = "metadata"
)

func (s EditMode) String() string {
	return string(s)
}

type PageSate string

const (
	Preview PageSate = "preview"
	Edit    PageSate = "edit"
	Normal  PageSate = "normal"
)

func (s PageSate) String() string {
	return string(s)
}
