package requests

import (
	"cozysystems.net/projects/CloudNative/repos/cozysystems.net/apiclient"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"fmt"
	"time"
)

func GetVMCatalogResourceID(apiC *apiclient.OnecloudAPICLient, vmName string ) (string, error){

	vm, err := GetVMCatalogResource(apiC, vmName)
	if err != nil {
		log.WithFields(log.Fields{"package": "requests","function": "GetVMCatalogResourceID",}).Debugf("Received error while trying to fetch Catalog resources for VM: %s Error -> %s", vmName, err)
		return "", err
	} else {
		log.WithFields(log.Fields{"package": "requests","function": "GetVMCatalogResourceID",}).Debugf("Returning Id for VM %s  ID : %s", vmName, vm.ID)
		return vm.ID, nil
	}
}

func GetVMCatalogResource(apiC *apiclient.OnecloudAPICLient, vmName string ) (*CatalogResource, error){

	path := "/catalog-service/api/consumer/resources/types/Infrastructure.Machine"
	log.WithFields(log.Fields{"package": "requests","function": "GetVMCatalogResource",}).Debugf("API Path: %s", path)
	getResourcesResp := new(GetResourcesOfARequestResponse)
	apiError := new(apiclient.APIError)
	log.WithFields(log.Fields{"package": "requests","function": "GetVMCatalogResource",}).Debugf("Applying Filter: name eq %s", vmName)
	params := &Params{Filter: "name eq '"+vmName+"'"}
	_,err := apiC.HTTCloudNativelient.New().Get(path).QueryStruct(params).
		Receive(getResourcesResp, apiError)

	if err != nil {
		log.WithFields(log.Fields{"package": "requests","function": "GetVMCatalogResource",}).Debugf("Received error while trying to fetch Catalog resources for VM: %s Error -> %s", vmName, err)
		return nil, err
	}
	if !(len(apiError.Errors) == 0){
		log.WithFields(log.Fields{"package": "requests","function": "GetVMCatalogResource",}).Debugf("API Error while trying to fetch Catalog resources for VM: %s Error -> %s", apiError)
		return nil, fmt.Errorf("%s", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
	}
	log.WithFields(log.Fields{"package": "requests","function": "GetVMCatalogResource",}).Debugf("Length of getResourcesResp.Content -> %s", len(getResourcesResp.Content))
	log.WithFields(log.Fields{"package": "requests","function": "GetVMCatalogResource",}).Debugf("Value of getResourcesResp.Content[0] -> %s", getResourcesResp.Content[0])
	return getResourcesResp.Content[0], nil
}

func (c *CatalogItemRequest) GetLatestStatus(apiC *apiclient.OnecloudAPICLient) (string, error)  {

	r, err := GetRequestResponse(c.ID, apiC)
	log.WithFields(log.Fields{"package": "requests","function": "GetLatestStatus",}).Debugf("Getting latest status for Catalog Item Request ID: %s", c.ID)
	apiError := new(apiclient.APIError)
	if err != nil {
		log.WithFields(log.Fields{"package": "requests","function": "GetLatestStatus",}).Debugf("Received error while trying to get latest status of Catalog request ID: %s Error -> %s", c.ID, err)
		return "", err
	}
	if !(len(apiError.Errors) == 0){
		log.WithFields(log.Fields{"package": "requests","function": "GetLatestStatus",}).Debugf("API Error while trying to get latest status of Catalog request ID: %s  Error -> %s", c.ID,apiError)
		return "", fmt.Errorf("%s", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
	}
	log.WithFields(log.Fields{"package": "requests","function": "GetLatestStatus",}).Debugf("Returning latest status of Catalog request ID: %s  Status -> %s", c.ID,r.Phase)
	return r.Phase, nil

}

func (c *CatalogItemRequest) WaitForCompletion(apiC *apiclient.OnecloudAPICLient) (string, error) {

	currentStatus,err := c.GetLatestStatus(apiC)
	log.WithFields(log.Fields{"package": "requests","function": "WaitForCompletion",}).Debugf("Current status for Catalog Item Request ID: %s Status -> %s", c.ID, currentStatus)
	for ok := true; ok; ok = !(currentStatus == SUCCESSFUL || currentStatus == FAILED || currentStatus == REJECTED) {
		currentStatus,err = c.GetLatestStatus(apiC)
		log.WithFields(log.Fields{"package": "requests","function": "WaitForCompletion",}).Debugf("Current status for Catalog Item Request ID: %s Status -> %s", c.ID, currentStatus)
		if err != nil {
			log.WithFields(log.Fields{"package": "requests","function": "WaitForCompletion",}).Debugf("Received error while trying to get latest status of Catalog request ID: %s Error -> %s", c.ID, err)
			return "", err
		}
		time.Sleep(15 * time.Second)
	}
	log.WithFields(log.Fields{"package": "requests","function": "WaitForCompletion",}).Debugf("Returning final status for Catalog Item Request ID: %s", currentStatus)
	return currentStatus,nil
}

func (c *CatalogItemRequest) GetCozyRequestIDFromCatalogItemRequest() (string, bool)  {

     entries := c.RequestData.Entries
     for _,kv := range entries{
     	if kv["key"] == "provider-*.Server.RequestId"{
     		t := kv["value"].(map[string]interface{})
     		return t["value"].(string),false
		}
	 }
	 return "", true
}

func (c *CatalogItemRequest) GetCatalogResourcesFromRequest(apiC *apiclient.OnecloudAPICLient, filter string) (*GetResourcesOfARequestResponse, error) {

	path := "/catalog-service/api/consumer/requests/"+c.ID+"/resources"
	log.WithFields(log.Fields{"package": "requests","function": "GetCatalogResourcesFromRequest",}).Debugf("API Path : %s", path)
	getResourcesResp := new(GetResourcesOfARequestResponse)
	apiError := new(apiclient.APIError)
	log.WithFields(log.Fields{"package": "requests","function": "GetCatalogResourcesFromRequest",}).Debugf("Applying filter on API call Filter: %s", filter)
	params := &Params{Filter: filter}
	_,err := apiC.HTTCloudNativelient.New().Get(path).QueryStruct(params).
		Receive(getResourcesResp, apiError)

	if err != nil {
		log.WithFields(log.Fields{"package": "requests","function": "GetCatalogResourcesFromRequest",}).Debugf("Received error while trying to get Catalog Resources from Catalog request ID: %s Error -> %s", c.ID, err)
		return nil, err
	}
	if !(len(apiError.Errors) == 0){
		log.WithFields(log.Fields{"package": "requests","function": "GetCatalogResourcesFromRequest",}).Debugf("API Error while trying to get Catalog Resources from Catalog request ID: %s  Error -> %s", c.ID,apiError)
		return nil, fmt.Errorf("%s", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
	}
	log.WithFields(log.Fields{"package": "requests","function": "GetCatalogResourcesFromRequest",}).Debugf("Successful API Response containing Catalog Resources from Catalog request ID: %s  API Response -> %s", c.ID,getResourcesResp)
	return getResourcesResp, nil
}

func (c *CatalogItemRequest) GetCatalogResourcesOfTypeVMFromReq(apiC *apiclient.OnecloudAPICLient, filter string) ([]*CatalogResource, error){

	log.WithFields(log.Fields{"package": "requests","function": "GetCatalogResourcesOfTypeVMFromReq",}).Debugf("Catalog request ID: %s  Filter: %s", c.ID,filter)
	var ctlgReqs []*CatalogResource
	x,err :=  c.GetCatalogResourcesFromRequest(apiC,filter)
	log.WithFields(log.Fields{"package": "requests","function": "GetCatalogResourcesOfTypeVMFromReq",}).Debugf("Response for  c.GetCatalogResourcesFromRequests: %s  Filter: %s", x,filter)
	apiError := new(apiclient.APIError)

	if err != nil {
		log.WithFields(log.Fields{"package": "requests","function": "GetCatalogResourcesOfTypeVMFromReq",}).Debugf("Received error while trying to GetCatalogResourcesOfTypeVMFromReq from Catalog request ID: %s Error -> %s", c.ID, err)
		return nil, err
	}
	if !(len(apiError.Errors) == 0){
		log.WithFields(log.Fields{"package": "requests","function": "GetCatalogResourcesOfTypeVMFromReq",}).Debugf("API Error while trying to  to GetCatalogResourcesOfTypeVMFromReq from Catalog request ID: %s  Error -> %s", c.ID,apiError)
		return nil, fmt.Errorf("%s", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
	}

	ctlgResources :=  x.Content

	for _,r := range ctlgResources {
		if x:= r.ResourceTypeRef.Label; x == "Virtual Machine" {
			ctlgReqs = append(ctlgReqs, r)
		}
	}
	log.WithFields(log.Fields{"package": "requests","function": "GetCatalogResourcesOfTypeVMFromReq",}).Debugf("Found %s Catalog resources of type VM from Catalog request ID: %s", len(ctlgReqs), c.ID)
	return  ctlgReqs,nil
}

func (c *CatalogResource) GetVirtualMachineName() string{
	log.WithFields(log.Fields{"package": "requests","function": "GetVirtualMachineName",}).Debugf("Returning VM Name for Catalog Resource ID: %s Virtual Machine Name: %s",c.ID, c.Name)
	return c.Name
}

func GetAllGenericRequests(apiC *apiclient.OnecloudAPICLient, filter string) (*ResponseForGetAllGenericRequests, error) {

	path := "/catalog-service/api/consumer/requests"
	log.WithFields(log.Fields{"package": "requests","function": "GetAllGenericRequests",}).Debugf("API Path : %s  Filter: %s", path, filter)
	getCatItemsResp := new(ResponseForGetAllGenericRequests)
	apiError := new(apiclient.APIError)
	jsonBody, jErr := json.Marshal(getCatItemsResp)

	if jErr != nil {
		log.WithFields(log.Fields{"package": "requests","function": "GetAllGenericRequests",}).Debugf("Error marshalling template as JSON. Error : %s", jErr)
		return nil, jErr
	} else {
		log.WithFields(log.Fields{"package": "requests","function": "GetAllGenericRequests",}).Debugf("JSON Request Info: %s", jsonBody)
	}

	// working fine params := &Params{Filter: "(dateCreated gt '2018-03-08T22:55:25.431Z' and dateCreated le '2018-03-08T23:14:59.431Z') and description eq 'Request from cozysystems.net'"}
	params := &Params{Filter: filter}
	_,err := apiC.HTTCloudNativelient.New().Get(path).QueryStruct(params).
		Receive(getCatItemsResp, apiError)

	if err != nil {
		log.WithFields(log.Fields{"package": "requests","function": "GetAllGenericRequests",}).Debugf("Received error while trying to GetAllGenericRequests based on filter: %s Error -> %s",  filter, err)
		return nil, err
	}
	if !(len(apiError.Errors) == 0){
		log.WithFields(log.Fields{"package": "requests","function": "GetAllGenericRequests",}).Debugf("API Error while trying to  to GetAllGenericRequests. Error is -> %s",apiError)
		return nil, fmt.Errorf("%s", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
	}
	log.WithFields(log.Fields{"package": "requests","function": "GetAllGenericRequests",}).Debugf("Successful API Response GetAllGenericRequests.  API Response -> %s", getCatItemsResp)
	return getCatItemsResp, nil


}

func FilterCatalogItemReqByCozyReqID(reqID string, c []*CatalogItemRequest) []*CatalogItemRequest{

	var ctlgReqs []*CatalogItemRequest
	for _,r := range c {
		if x,_ := r.GetCozyRequestIDFromCatalogItemRequest(); x == reqID {
			ctlgReqs = append(ctlgReqs, r)
		}
	}
	return  ctlgReqs
}

// Keeping startDate/endDate as strings. Need to check to time.Time and come up with a way to assert this user input properly

func GetAllCatalogItemRequests(user string, startDate string, endDate string, apiC *apiclient.OnecloudAPICLient) ([]*CatalogItemRequest, error){

	filter := "(dateCreated gt '"+startDate+"' and dateCreated le '"+endDate+"') and reasons eq 'Standard VM Request' and substringof('"+user+"',tolower(description))"
	resp,err := GetAllGenericRequests(apiC, filter)
	apiError := new(apiclient.APIError)
	if err != nil {
		return nil, err
	}
	if !(len(apiError.Errors) == 0){
		fmt.Println(apiError)
		return nil, fmt.Errorf("%s", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
	}

	return resp.Content,nil
}

func GetRequestResponse(reqId string, apiC *apiclient.OnecloudAPICLient) (*CatalogItemRequest, error)  {

	path := "/catalog-service/api/consumer/requests/"+reqId+""
	reqResp := new(CatalogItemRequest)
	apiError := new(apiclient.APIError)
	_,err := apiC.HTTCloudNativelient.New().Get(path).
		Receive(reqResp, apiError)
	if err != nil {
		return nil, err
	}
	if !(len(apiError.Errors) == 0){
		fmt.Println(apiError)
		return nil, fmt.Errorf("%s", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
	}
	return reqResp, nil
}