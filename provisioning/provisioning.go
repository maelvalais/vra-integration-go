package provisioning

import (
	"cozysystems.net/projects/CloudNative/repos/cozysystems.net/apiclient"
	"fmt"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)


func RequestVMfromTemplate(apiC *apiclient.OnecloudAPICLient, template *OnecloudVMRequestTemplate) (*OnecloudVMRequestResponse, error){
	path := fmt.Sprintf("/catalog-service/api/consumer/entitledCatalogItems/%s"+
		"/requests", template.CatalogItemID)
	log.WithFields(log.Fields{"package": "provisioning","function": "RequestVMfromTemplate",}).Debugf("API Path: %s", path)
	requestMachineRes := new(OnecloudVMRequestResponse)
	apiError := new(apiclient.APIError)
	jsonBody, jErr := json.Marshal(template)
	if jErr != nil {
		log.WithFields(log.Fields{"package": "provisioning","function": "RequestVMfromTemplate",}).Debugf("Error marshalling template as JSON")
		return nil, jErr
	} else {
		log.WithFields(log.Fields{"package": "provisioning","function": "RequestVMfromTemplate",}).Debugf("JSON Request Info: %s", jsonBody)
	}
	_,err := apiC.HTTCloudNativelient.New().Post(path).BodyJSON(template).
		     Receive(requestMachineRes, apiError)
	if err != nil {
		return nil, err
	}
	if !(len(apiError.Errors) == 0){
		return nil, fmt.Errorf("%s", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
	}
	return requestMachineRes, nil

}


