package provisioning

import (
	"time"
)

type OnecloudResourceCfgData struct {
	Product     string  `json:"product"`
	Clustered   bool    `json:"clustered"`
	Cpu         float32 `json:"cpu"`
	Datacenter  string  `json:"datacenter"`
	Environment string  `json:"environment"`
	Instances   float32 `json:"instances"`
	Memory      float32 `json:"memory"`
	Network     string  `json:"network"`
	Disk1Mount  string  `json:"disk1Mount"`
	Disk1Size   int     `json:"disk1Size"`
	NumDisks    float32 `json:"numDisks"`
	OS          string  `json:"os"`
	Role        string  `json:"role"`
	Type        string  `json:"type"`
}

type OnecloudVMRequestTemplate struct {
	BusinessGroupID string `json:"businessGroupId"`
	CatalogItemID   string `json:"catalogItemId"`
	Data *OnecloudResourceCfgData    `json:"data"`
	Description     string  `json:"description"`
	RequestedFor    string  `json:"requestedFor"`
	Type            string  `json:"type"`
}

type OnecloudVMRequestResponse struct {
	ID           string      `json:"id"`
	IconID       string      `json:"iconId"`
	Version      int         `json:"version"`
	State        string      `json:"state"`
	Description  string      `json:"description"`
	Reasons      interface{} `json:"reasons"`
	RequestedFor string      `json:"requestedFor"`
	RequestedBy  string      `json:"requestedBy"`
	Organization struct {
		TenantRef      string `json:"tenantRef"`
		TenantLabel    string `json:"tenantLabel"`
		SubtenantRef   string `json:"subtenantRef"`
		SubtenantLabel string `json:"subtenantLabel"`
	} `json:"organization"`

	RequestorEntitlementID   string                 `json:"requestorEntitlementId"`
	PreApprovalID            string                 `json:"preApprovalId"`
	PostApprovalID           string                 `json:"postApprovalId"`
	DateCreated              time.Time              `json:"dateCreated"`
	LastUpdated              time.Time              `json:"lastUpdated"`
	DateSubmitted            time.Time              `json:"dateSubmitted"`
	DateApproved             time.Time              `json:"dateApproved"`
	DateCompleted            time.Time              `json:"dateCompleted"`
	Quote                    interface{}            `json:"quote"`
	RequestData              map[string]interface{} `json:"requestData"`
	RequestCompletion        string                 `json:"requestCompletion"`
	RetriesRemaining         int                    `json:"retriesRemaining"`
	RequestedItemName        string                 `json:"requestedItemName"`
	RequestedItemDescription string                 `json:"requestedItemDescription"`
	Components               string                 `json:"components"`
	StateName                string                 `json:"stateName"`

	CatalogItemProviderBinding struct {
		BindingID   string `json:"bindingId"`
		ProviderRef struct {
			ID    string `json:"id"`
			Label string `json:"label"`
		} `json:"providerRef"`
	} `json:"catalogItemProviderBinding"`

	Phase           string `json:"phase"`
	ApprovalStatus  string `json:"approvalStatus"`
	ExecutionStatus string `json:"executionStatus"`
	WaitingStatus   string `json:"waitingStatus"`
	CatalogItemRef  struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	} `json:"catalogItemRef"`
}
