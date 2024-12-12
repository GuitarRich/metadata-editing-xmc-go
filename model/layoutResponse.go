package model

type RouteData struct {
	Name         string                            `json:"name"`
	DisplayName  string                            `json:"displayName"`
	Fields       map[string]interface{}            `json:"fields"`
	DatabaseName string                            `json:"databaseName"`
	DeviceID     string                            `json:"deviceId"`
	ItemID       string                            `json:"itemId"`
	ItemLanguage string                            `json:"itemLanguage"`
	ItemVersion  int                               `json:"itemVersion"`
	LayoutID     string                            `json:"layoutId"`
	TemplateID   string                            `json:"templateId"`
	TemplateName string                            `json:"templateName"`
	Placeholders map[string][]PlaceholderComponent `json:"placeholders"`
}

type Rendered struct {
	Sitecore Sitecore `json:"sitecore"`
}

type Sitecore struct {
	Context SitecoreContext `json:"context"`
	Route   RouteData       `json:"route"`
}

type SitecoreContext struct {
	PageEditing bool `json:"pageEditing"`
	Site        struct {
		Name string `json:"name"`
	} `json:"site"`
	PageState  string `json:"pageState"`
	EditMode   string `json:"editMode"`
	Language   string `json:"language"`
	ItemPath   string `json:"itemPath"`
	ClientData struct {
		HorizonCanvasState struct {
			ItemId      string  `json:"itemId"`
			ItemVersion float64 `json:"itemVersion"`
			SiteName    string  `json:"siteName"`
			Language    string  `json:"language"`
			DeviceId    string  `json:"deviceId"`
			PageMode    string  `json:"pageMode"`
			Variant     string  `json:"variant"`
		} `json:"hrz-canvas-state"`
		HorizonCanvasVerificationToken string `json:"hrz-canvas-verification-token"`
	} `json:"clientData"`
	ClientScripts []string `json:"clientScripts"`
}

type LayoutResponse struct {
	Data struct {
		Layout struct {
			Item struct {
				Rendered Rendered `json:"rendered"`
			} `json:"item"`
		} `json:"layout"`
	} `json:"data"`
}

type EditingResponse struct {
	Data struct {
		Item struct {
			Rendered Rendered `json:"rendered"`
		} `json:"item"`
	} `json:"data"`
}

type NotFoundPageResponse struct {
	Data struct {
		Site struct {
			SiteInfo struct {
				ErrorHandling struct {
					NotFoundPage struct {
						Rendered Rendered `json:"rendered"`
					} `json:"notFoundPage"`
				} `json:"errorHandling"`
			} `json:"siteInfo"`
		} `json:"site"`
	} `json:"data"`
}
