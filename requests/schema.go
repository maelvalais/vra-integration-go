package requests

import (
	"time"
)

type Params struct {
	Filter string `url:"$filter,omitempty"`
}

type CatalogResource struct {

	Type                 string `json:"@type"`
	ID                   string `json:"id"`
	IconID               string `json:"iconId"`
	ResourceTypeRef struct{
		ID        string `json:"id"`
		Label     string `json:"label"`
	} `json:"resourceTypeRef"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Status               string `json:"status"`
	CatalogItem     struct{
		ID        string `json:"id"`
		Label     string `json:"label"`
	} `json:"catalogItem"`
	RequestId            string `json:"requestId"`
	RequestState         string `json:"requestState"`
	ProviderBinding  struct{
		BindingId   string `json:"bindingId"`
		ProviderRef struct{
			ID        string `json:"id"`
			Label     string `json:"label"`
		} `json:"providerRef"`
	} `json:"providerBinding"`
	Owners [] struct{
		TenantName string `json:"tenantName"`
		Ref        string `json:"ref"`
		Type       string `json:"type"`
		Value      string `json:"value"`
	} `json:"owners"`
	Organization struct {
		TenantRef      string `json:"tenantRef"`
		TenantLabel    string `json:"tenantLabel"`
		SubtenantRef   string `json:"subtenantRef"`
		SubtenantLabel string `json:"subtenantLabel"`
	} `json:"organization"`
	DateCreated        string `json:"dateCreated"`
	LastUpdated        string `json:"lastUpdated"`
	HasLease           bool   `json:"hasLease"`
	Lease struct{
		Start    string `json:"start"`
		End      string `json:"end"`
	} `json:"lease"`
	LeaseForDisplay    string `json:"leaseForDisplay"`
	HasCosts           bool   `json:"hasCosts"`
	Costs              string `json:"costs"`
	CostToDate		   string `json:"costToDate"`
	TotalCost		   string `json:"totalCost"`
	ExpenseMonthToDate string `json:"expenseMonthToDate"`
	ParentResourceRef  interface{} `json:"parentResourceRef"`
	HasChildren		   bool   `json:"hasChildren"`
	Operations [] struct{
		Name           string `json:"name"`
		Description    string `json:"description"`
		IconId         string `json:"iconId"`
		Type           string `json:"type"`
		Id             string `json:"id"`
		ExtensionId    string `json:"extensionId"`
		ProviderTypeId string `json:"providerTypeId"`
		BindingId      string `json:"bindingId"`
		HasForm        bool   `json:"hasForm"`
		FormScale      string `json:"formScale"`
	} `json:"operations"`
	Forms struct{
		CatalogResourceInfoHidden  bool `json:"catalogResourceInfoHidden"`
		Details struct{
			Type    string `json:"type"`
			FormId  string `json:"formId"`
		} `json:"details"`
	} `json:"forms"`
	ResourceData struct{
		Entries []interface{} `json:"entries"`
	} `json:"resourceData"`
	DestroyDate      string   `json:"destroyDate"`
	PendingRequests  []interface{}
}

type GetResourcesOfARequestResponse struct {
	Links    []interface{}         `json:"links"`
	Content  []*CatalogResource    `json:"content"`
	Metadata struct{
		Size          int  `json:"size"`
		TotalElements int  `json:"totalElements"`
		TotalPages    int  `json:"totalPages"`
		Number        int  `json:"number"`
		Offset        int  `json:"offset"`
	}  `json:"metadata"`
}

type CatalogItemRequest struct {

	Type           string `json:"@type"`
	ID             string `json:"id"`
	IconID         string `json:"iconId"`
	Version        int    `json:"version"`
	RequestNumber  int    `json:"requestNumber"`
	State          string `json:"state"`
	Description    string `json:"description"`
	Reasons        string `json:"reasons"`
	RequestedFor   string `json:"requestedFor"`
	RequestedBy    string `json:"requestedBy"`
	Organization struct {
		TenantRef      string `json:"tenantRef"`
		TenantLabel    string `json:"tenantLabel"`
		SubtenantRef   string `json:"subtenantRef"`
		SubtenantLabel string `json:"subtenantLabel"`
	} `json:"organization"`
    RequestorEntitlementID string     `json:"requestorEntitlementId"`
    PreapprovalID          string     `json:"preApprovalId"`
    PostapprovalID         string     `json:"postApprovalId"`
    DateCreated            time.Time  `json:"dateCreated"`
    LastUpdated            time.Time  `json:"lastUpdated"`
	DateSubmitted          time.Time  `json:"dateSubmitted"`
	DateApproved           time.Time  `json:"dateApproved"`
	DateCompleted          time.Time  `json:"dateCompleted"`
	Quote struct{
		LeasePeriod        string     `json:"leasePeriod"`
		LeaseRate          string     `json:"leaseRate"`
		TotalLeaseCost     string     `json:"totalLeaseCost"`
	}
	RequestCompletion struct{
		RequestCompletionState string  `json:"requestCompletionState"`
		CompletionDetails      string  `json:"completionDetails"`
		ResourceBindingIds     string  `json:"resourceBindingIds"`
	}
	//RequestData                map[string] interface{} `json:"requestData"`
	RequestData  struct{
		Entries []map[string] interface{} `json:"entries"`
	} `json:"requestData"`
	RetriesRemaining           int     `json:"retriesRemaining"`
	RequestedItemName          string  `json:"requestedItemName"`
	RequestedItemDescription   string  `json:"requestedItemDescription"`
	Components                 string  `json:"components"`
	StateName                  string  `json:"stateName"`
	CatalogItemRef struct{
		ID          string  `json:"id"`
		Label       string  `json:"label"`
	}  `json:"catalogItemRef"`
	CatalogItemProviderBinding struct{
		BindingId      string   `json:"bindingId"`
		ProviderRef struct{
			ID          string  `json:"id"`
			Label       string  `json:"label"`
		} `json:"providerRef"`

	} `json:"catalogItemProviderBinding"`
	Phase             string `json:"phase"`
	ExecutionStatus   string `json:"executionStatus"`
	WaitingStatus     string `json:"waitingStatus"`
	ApprovalStatus    string `json:"ApprovalStatus"`
}


type ResponseForGetAllGenericRequests struct {
	Links    []interface{}         `json:"links"`
	Content  []*CatalogItemRequest  `json:"content"`
	Metadata struct{
		Size          int  `json:"size"`
		TotalElements int  `json:"totalElements"`
		TotalPages    int  `json:"totalPages"`
		Number        int  `json:"number"`
		Offset        int  `json:"offset"`
	}  `json:"metadata"`
}