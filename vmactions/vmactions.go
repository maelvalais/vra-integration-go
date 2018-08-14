package vmactions

import (
	"cozysystems.net/projects/CloudNative/repos/cozysystems.net/apiclient"
	"fmt"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func (vmActs *VMActions) GetConnectUsingSSHAction() (*ConsumerResourceOperation){

	for _,x := range vmActs.Content {
		if x.Name == CONNNECT_USING_SSH {
			return  x
		}
	}
	return nil
}

func (vmActs *VMActions) GetCreateSnapshotAction() (*ConsumerResourceOperation){

	for _,x := range vmActs.Content {
		if x.Name == CREATE_SNAPSHOT {
			return  x
		}
	}
	return nil
}

func (vmActs *VMActions) GetDecommissionAction() (*ConsumerResourceOperation){

	for _,x := range vmActs.Content {
		if x.Name == DECOMMISSION {
			return  x
		}
	}
	return nil
}

func (vmActs *VMActions) GetForceDecommissionAction() (*ConsumerResourceOperation){

	for _,x := range vmActs.Content {
		if x.Name == FORCE_DECOMMISSION {
			return  x
		}
	}
	return nil
}

func (vmActs *VMActions) GetHardRebootVMAction() (*ConsumerResourceOperation){

	for _,x := range vmActs.Content {
		if x.Name == HARD_REBOOT_VM {
			return  x
		}
	}
	return nil
}

func (vmActs *VMActions) GetHardShutdownVMAction() (*ConsumerResourceOperation){

	for _,x := range vmActs.Content {
		if x.Name == HARD_SHUTDOWN_VM {
			return  x
		}
	}
	return nil
}

func (vmActs *VMActions) GetRevertSnapshotVMAction() (*ConsumerResourceOperation){

	for _,x := range vmActs.Content {
		if x.Name == REVERT_SNAPSHOT {
			return  x
		}
	}
	return nil
}

func (vmActs *VMActions) GetSoftRebootVMAction() (*ConsumerResourceOperation){

	for _,x := range vmActs.Content {
		if x.Name == SOFT_REBOOT_VM {
			return  x
		}
	}
	return nil
}

func (vmActs *VMActions) GetSoftShutdownVMAction() (*ConsumerResourceOperation){

	for _,x := range vmActs.Content {
		if x.Name == SOFT_SHUTDOWN_VM {
			return  x
		}
	}
	return nil
}

func (vmActs *VMActions) ExecuteConnectUsingSSHAction(vmID string, apiC *apiclient.OnecloudAPICLient) (error){

	action := vmActs.GetConnectUsingSSHAction()
	if action != nil {
		err := action.Execute(vmID, apiC)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("Operation not found on the VM. Access to this action is not available")
}

func (vmActs *VMActions) ExecuteCreateSnapshotAction(vmID string, apiC *apiclient.OnecloudAPICLient) (error){

	action := vmActs.GetCreateSnapshotAction()
	if action != nil {
		err := action.Execute(vmID, apiC)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("Operation not found on the VM. Access to this action is not available")
}

func (vmActs *VMActions) ExecuteDecommissionAction(vmID string, apiC *apiclient.OnecloudAPICLient) (error){

	action := vmActs.GetDecommissionAction()
	if action != nil {
		err := action.Execute(vmID, apiC)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("Operation not found on the VM. Access to this action is not available")
}

func (vmActs *VMActions) ExecuteForceDecommissionAction(vmID string, apiC *apiclient.OnecloudAPICLient) (error){

	action := vmActs.GetForceDecommissionAction()
	if action != nil {
		err := action.Execute(vmID, apiC)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("Operation not found on the VM. Access to this action is not available")
}

func (vmActs *VMActions) ExecuteHardRebootVMAction(vmID string, apiC *apiclient.OnecloudAPICLient) (error){

	action := vmActs.GetHardRebootVMAction()
	if action != nil {
		err := action.Execute(vmID, apiC)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("Operation not found on the VM. Access to this action is not available")
}

func (vmActs *VMActions) ExecuteHardShutdownVMAction(vmID string, apiC *apiclient.OnecloudAPICLient) (error){

	action := vmActs.GetHardShutdownVMAction()
	if action != nil {
		err := action.Execute(vmID, apiC)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("Operation not found on the VM. Access to this action is not available")
}

func (vmActs *VMActions) ExecuteRevertSnapshotVMAction(vmID string, apiC *apiclient.OnecloudAPICLient) (error){

	action := vmActs.GetRevertSnapshotVMAction()
	if action != nil {
		err := action.Execute(vmID, apiC)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("Operation not found on the VM. Access to this action is not available")
}

func (vmActs *VMActions) ExecuteSoftRebootVMAction(vmID string, apiC *apiclient.OnecloudAPICLient) (error){

	action := vmActs.GetSoftRebootVMAction()
	if action != nil {
		err := action.Execute(vmID, apiC)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("Operation not found on the VM. Access to this action is not available")
}

func (vmActs *VMActions) ExecuteSoftShutdownVMAction(vmID string, apiC *apiclient.OnecloudAPICLient) (error){

	action := vmActs.GetSoftShutdownVMAction()
	if action != nil {
		err := action.Execute(vmID, apiC)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return fmt.Errorf("Operation not found on the VM. Access to this action is not available")
}


func (c *ConsumerResourceOperation) Execute(vmID string, apiC *apiclient.OnecloudAPICLient) (error){

	path := "/catalog-service/api/consumer/resources/"+vmID+"/actions/"+c.ID+"/requests"
	vmActionTemplate, err := c.GetActionTemplate(apiC, vmID)
	apiError := new(apiclient.APIError)
	jsonBody, jErr := json.Marshal(vmActionTemplate)
	if jErr != nil {
		log.Printf("Error marshalling template as JSON")
		return jErr
	} else {
		log.Printf("JSON Request Info: %s", jsonBody)
	}
	if err != nil {
		return err
	}
	if !(len(apiError.Errors) == 0){
		fmt.Println(apiError)
		return fmt.Errorf("%s", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
	}

	resp,err := apiC.HTTCloudNativelient.New().Post(path).BodyJSON(vmActionTemplate).Receive(nil, nil)

	if err != nil {
		return err
	}
	if !(len(apiError.Errors) == 0){
		return fmt.Errorf("%s", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
	}
	if resp.StatusCode != 201 {
		return fmt.Errorf("expected %d, got %d", 201, resp.StatusCode)
	}
	return nil

}

func GetVMActions(vmID string, apiC *apiclient.OnecloudAPICLient) (*VMActions, error) {

	path := "/catalog-service/api/consumer/resources/"+vmID+"/actions/"
	getVMActionsResp := new(VMActions)
	apiError := new(apiclient.APIError)
	_,err := apiC.HTTCloudNativelient.New().Get(path).Receive(getVMActionsResp, apiError)

	if err != nil {
		return nil, err
	}
	if !(len(apiError.Errors) == 0){
		fmt.Println(apiError)
		return nil, fmt.Errorf("%s", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
	}
	return getVMActionsResp, nil
}

func (c *ConsumerResourceOperation) GetActionTemplate(apiC *apiclient.OnecloudAPICLient, vmID string) (*VMActionTemplate, error) {

	path := "/catalog-service/api/consumer/resources/"+vmID+"/actions/"+c.ID+"/requests/template"
	getTemplateResp := new(VMActionTemplate)
	apiError := new(apiclient.APIError)
	_,err := apiC.HTTCloudNativelient.New().Get(path).
		Receive(getTemplateResp, apiError)

	if err != nil {
		return nil, err
	}
	if !(len(apiError.Errors) == 0){
		fmt.Println(apiError)
		return nil, fmt.Errorf("%s", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
	}

	// This is a hack to workaround a bug in onecloud where decommission action has have a flag set to false so that VMS are imm destroyed instead of being archived for 7 days.
	if getTemplateResp.Data.ProviderArchiveFlag {
		getTemplateResp.Data.ProviderArchiveFlag = false
	}
	return getTemplateResp, nil
}

