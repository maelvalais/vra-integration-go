package vmactions


type ConsumerResourceOperation struct {
	Name           string `json:"name"`
	Description    string `description`
	IconId         string `json:"iconId"`
	Type           string `json:"type"`
	ID             string `json:"id"`
	ExtensionId    string `json:"extensionId"`
	ProviderTypeId string `json:"providerTypeId"`
	BindingId      string `json:"bindingId"`
	HasForm        bool   `json:"hasForm"`
	FormScale      string `json:"formScale"`

}

type VMActions struct {
	Links    []interface{}                   `json:"links"`
	Content  []*ConsumerResourceOperation    `json:"content"`

}

type VMActionTemplate struct {
	Type         string           `json:"type"`
	ResourceId   string           `json:"resourceId"`
	ActionId     string           `json:"actionId"`
	Description  string           `json:"description"`
	//Data  map[string] interface{} `json:"data"`
	Data  struct{
		ProviderASD_PRESENTATION_INSTANCE *string  `json:"provider-__ASD_PRESENTATION_INSTANCE"`
		ProviderArchiveFlag               bool    `json:"provider-archiveFlag"`
		ProviderCO                        *string  `json:"provider-co"`
		ProviderDecommSched               *string  `json:"provider-decommSched"`
	} `json:"data"`
}