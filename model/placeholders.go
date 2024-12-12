package model

type PlaceholderComponent struct {
	UID           string `json:"uid"`
	ComponentName string `json:"componentName"`
	DataSource    string `json:"dataSource"`
	Params        struct {
		GridParameters       string `json:"GridParameters"`
		FieldNames           string `json:"FieldNames"`
		Styles               string `json:"Styles"`
		DynamicPlaceholderID string `json:"DynamicPlaceholderId"`
		Sig                  string `json:"sig"`
		Ph                   string `json:"ph"`
		BackgroundImage      string `json:"BackgroundImage"`
	} `json:"params"`
	Fields       interface{}                       `json:"fields"`
	Placeholders map[string][]PlaceholderComponent `json:"placeholders"`
}

type Placeholder struct {
	Components []PlaceholderComponent `json:"components"`
}
