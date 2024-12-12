package model

type Datasource struct {
	Id       string  `json:"id"`
	Language string  `json:"language"`
	Revision string  `json:"revision"`
	Version  float64 `json:"version"`
}

type MetadataData struct {
	Datasource Datasource `json:"datasource"`
	Title      string     `json:"title"`
	FieldId    string     `json:"fieldId"`
	FieldType  string     `json:"fieldType"`
	RawValue   string     `json:"rawValue"`
}

type RichTextField struct {
	Value    interface{}  `json:"value"`
	Editable string       `json:"editable"`
	Metadata MetadataData `json:"metadata"`
}

type ImageField struct {
	Value struct {
		Src    string `json:"src"`
		Alt    string `json:"alt"`
		Width  string `json:"width"`
		Height string `json:"height"`
	}
	Metadata MetadataData `json:"metadata"`
}

type LinkField struct {
	Value struct {
		Href        string `json:"href"`
		Anchor      string `json:"anchor"`
		Querystring string `json:"querystring"`
		Text        string `json:"text"`
		Target      string `json:"target"`
		Title       string `json:"title"`
		Class       string `json:"class"`
	} `json:"value"`
	Metadata MetadataData `json:"metadata"`
}

type ScField struct {
	Value    interface{}  `json:"value"`
	Metadata MetadataData `json:"metadata"`
}

type ImageFragment struct {
	Src    string `xml:"src,attr"`
	Alt    string `xml:"alt,attr"`
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
}
