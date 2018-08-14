package main

import (
	"fmt"
	"cozysystems.net/projects/CloudNative/repos/cozysystems.net/apiclient"
	"cozysystems.net/projects/CloudNative/repos/cozysystems.net/commonutils"
	"cozysystems.net/projects/CloudNative/repos/cozysystems.net/provisioning"
	"cozysystems.net/projects/CloudNative/repos/cozysystems.net/requests"
	"cozysystems.net/projects/CloudNative/repos/cozysystems.net/vmactions"
	strftime "github.com/jehiah/go-strftime"
	"github.com/magiconair/properties"
	log "github.com/sirupsen/logrus"
	"flag"
	"os"
	"strings"
)

func main() {

	log.SetLevel(log.DebugLevel)
	fmt.Println("************************************************ USAGE ************************************************")
	fmt.Println("cozysystems.net -username=pudupatr -password=*** -tenant=adp -action=createVMs  -vmtemplate=/home/rajan/vm.properties")
	fmt.Println("cozysystems.net -username=pudupatr -password=*** -tenant=adp -action=destroyVMs -vmlist=CDLCLDVMGXXYY,CDLCLDVMGXXYY,CDLCLDVMGXXYY")
	fmt.Println("")
	fmt.Println("Sample VM template file can be found at https://cozysystems.net/projects/CloudNative/repos/cozysystems.net/browse/examples/vm.properties")
	fmt.Println("*******************************************************************************************************")

	vmTemplateFile := flag.String("vmtemplate","","VMTemplate file supplied as input")
	action := flag.String("action","","Action requested by User")
	vmLst := flag.String("vmlist","","List of VMs to be destroyed")
	username := flag.String("username","","Onecloud Username")
	password := flag.String("password","","Onecloud Password")
	tenant := flag.String("tenant","","Onecloud Tenant")
	flag.Parse()
	vmTemplate := *vmTemplateFile
	act := *action
	vms := *vmLst
	user := *username
	passwd := *password
	ten := *tenant

	oneCloudClient,err := apiclient.NewAPIClient(user,passwd,ten,"smartcloud.es.ad.adp.com","443")
	if err != nil {
		log.WithFields(log.Fields{"package": "main","function": "main",}).Fatal("OneCloud Client creation failed . Fatal Error: %s", err)
		os.Exit(1)
	}
	log.WithFields(log.Fields{"package": "main","function": "main",}).Debugf("OneCloud client AUTH Token : %s", oneCloudClient.AuthToken)


	if act == "createVMs"{
		p := properties.MustLoadFile(vmTemplate, properties.UTF8)

		type Config struct {
			BusinessGroupID string    `properties:"businessGroupID"`
			CatalogItemID   string    `properties:"catalogItemID"`
			Product         string    `properties:"product"`
			Clustered       bool      `properties:"clustered,default=false"`
			Cpu             float32   `properties:"cpu"`
			Datacenter      string    `properties:"datacenter"`
			Environment     string    `properties:"environment"`
			Instances       float32   `properties:"instances"`
			Memory          float32   `properties:"memory"`
			Network         string    `properties:"network"`
			Disk1Mount      string    `properties:"disk1Mount"`
			Disk1Size       int       `properties:"disk1Size"`
			NumDisks        float32   `properties:"numDisks"`
			OS              string    `properties:"os"`
			Role            string    `properties:"role"`
			Type            string    `properties:"type"`
		}
		var cfg Config
		if err := p.Decode(&cfg); err != nil {
			log.Fatal(err)
		}

		data := &provisioning.OnecloudResourceCfgData{
			Product: strings.TrimSpace(cfg.Product),
			Clustered: false,
			Cpu:  cfg.Cpu,
			Datacenter: strings.TrimSpace(cfg.Datacenter),
			Environment: strings.TrimSpace(cfg.Environment),
			Instances: cfg.Instances,
			Memory: cfg.Memory,
			Network: cfg.Network,
			Disk1Mount: cfg.Disk1Mount,
			Disk1Size: cfg.Disk1Size,
			NumDisks: cfg.NumDisks,
			OS: strings.TrimSpace(cfg.OS),
			Role: strings.TrimSpace(cfg.Role),
			Type: strings.TrimSpace(cfg.Type),
		}

		reqTemplate := &provisioning.OnecloudVMRequestTemplate{
			BusinessGroupID: strings.TrimSpace(cfg.BusinessGroupID),
			CatalogItemID: strings.TrimSpace(cfg.CatalogItemID),
			Data: data,
			Description: "Request from CaaS teams' Onecloud GO library",
			RequestedFor: user+"@ES.AD.ADP.COM",
			Type: "com.vmware.vcac.catalog.domain.request.CatalogItemProvisioningRequest",
		}

		log.WithFields(log.Fields{"package": "main","function": "main",}).Debugf("OnecloudResourceCfgData: %s", data)
		log.WithFields(log.Fields{"package": "main","function": "main",}).Debugf("OnecloudVMRequestTemplate: %s", reqTemplate)

		log.WithFields(log.Fields{"package": "main","function": "main",}).Debugf("User has requested for -action=createVMs")
		strArr := CreateVMS(oneCloudClient,reqTemplate)
		log.WithFields(log.Fields{"package": "main","function": "main",}).Debugf("Created VMs. This is the List : %s", strArr)
		if int(reqTemplate.Data.Instances) != len(strArr) {
			log.WithFields(log.Fields{"package": "main","function": "main",}).Debugf("Requested No of VMs : %s | VMs created : %s", int(reqTemplate.Data.Instances), len(strArr))
			log.WithFields(log.Fields{"package": "main","function": "main",}).Fatal("OneCloud FAILED to create the requested number of VMs. Created VMs (%s) < Requested VMs (%s)",len(strArr),int(reqTemplate.Data.Instances))
		}
	} else if act == "destroyVMs"{
		log.WithFields(log.Fields{"package": "main","function": "main",}).Debugf("User has requested for -action=destroyVMs VM list -> %s",vms)
		for _,vm := range strings.Split(vms,","){
			vm = strings.TrimSpace(vm)
			DestroyVMs(oneCloudClient, vm)
		}
		log.WithFields(log.Fields{"package": "main","function": "main",}).Debugf("Successfully destroyed the following  VMs-> %s",vms)
	}

}

func CreateVMS(oneCloudClient *apiclient.OnecloudAPICLient, reqTemplate *provisioning.OnecloudVMRequestTemplate) []string {

	OnecloudVMRequestResponse,_ := provisioning.RequestVMfromTemplate(oneCloudClient, reqTemplate)
	log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("VRA response for VMCreation request -> %s", OnecloudVMRequestResponse)
	adpReqID := OnecloudVMRequestResponse.ID
	log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("VRA ADP Request ID : %s", adpReqID)
	r,_ := requests.GetRequestResponse(adpReqID,oneCloudClient)
	currentStatus,err := r.GetLatestStatus(oneCloudClient)
	log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("ADP Request ID status : %s", currentStatus)
	if err == nil {
		r.WaitForCompletion(oneCloudClient)

	} else {
		log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("ADP Request ID %s has failed! Fatal Error!", adpReqID)
		log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Fatal(err)
	}

	startDate := strftime.Format("%Y-%m-%dT%H:%M:%S.000Z", OnecloudVMRequestResponse.DateSubmitted.UTC())
	endDate   := commonutils.Add1HourToDateString(startDate)
	log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("Filtering VRA requests based on startDate: %s & endDate: %s", startDate, endDate)

	//CatalogItemReqs, err := requests.GetAllCatalogItemRequests("2018-03-12T19:37:25.431Z","2018-03-12T20:32:25.431Z",oneCloudClient)
	user := strings.Split(reqTemplate.RequestedFor,"@")[0]
	log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("User requesting the VMS is -> : %s", user)
	CatalogItemReqs, err := requests.GetAllCatalogItemRequests(user,startDate,endDate,oneCloudClient)
	if err == nil {
		log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("Successfully filtered CatalogItemRequests based on start & end date. No of requests found: %s", len(CatalogItemReqs))
	} else {
		log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("Unable to filter CatalogItemRequests based on start & end date! Fatal Error!")
		log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Fatal(err)
	}


	reqs := requests.FilterCatalogItemReqByADPReqID(adpReqID,CatalogItemReqs)
	log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("Successfully filtered CatalogItemRequests based on ADP request id: %s.  No of requests found: %s",adpReqID, len(CatalogItemReqs))

	var vmList []string

	for _,req := range reqs{
		log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("VRA internal request status: %s",req.ID)
		str, err := req.WaitForCompletion(oneCloudClient)
		if err != nil {
			log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("VRA internal Request ID %s has failed! Fatal Error! STATUS: %s", req.ID, str)
			log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Fatal(err)
		} else {
			log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("VRA internal Request ID %s is SUCCESSFUL! Proceeding further. Received STATUS: %s", req.ID, str)
		}
		log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("Getting CatalogResources of the VRA request id: %s",req.ID)
		resp, err := req.GetCatalogResourcesOfTypeVMFromReq(oneCloudClient,"")
		log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("No of virtual machines for request: %s is %s",req.ID, len(resp))
		if err == nil {
			for _,y := range resp {
				vmList = append(vmList, y.GetVirtualMachineName())
				log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("Request id: %s  Virtual Machine Name: %s",req.ID, y.GetVirtualMachineName())
			}

		} else {
			log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("Unable to get catalog resources for the request ID: %s! Fatal Error!", req.ID)
			log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Fatal(err)
		}
	}

	log.WithFields(log.Fields{"package": "main","function": "CreateVMS",}).Debugf("Final Virtual Machine List: %s",vmList)

	f, err := os.Create("vmlist.file")
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range vmList {
		fmt.Fprintln(f, value)
	}
	return  vmList
}


func DestroyVMs(oneCloudClient *apiclient.OnecloudAPICLient, vmName string) {
	log.WithFields(log.Fields{"package": "main","function": "DestroyVMs",}).Debugf("Destroy VM request received. Virtual machine name: %s", vmName)
	catalogResID, err := requests.GetVMCatalogResourceID(oneCloudClient,vmName)
	log.WithFields(log.Fields{"package": "main","function": "DestroyVMs",}).Debugf("Virtual machine ID: %s", catalogResID)
	if err != nil {
		log.WithFields(log.Fields{"package": "main","function": "DestroyVMs",}).Debugf("Unable to get destroy VM. Virtual machine name: %s! Fatal Error!", vmName)
		log.WithFields(log.Fields{"package": "main","function": "DestroyVMs",}).Fatal(err)
	}
	vmactions, err := vmactions.GetVMActions(catalogResID, oneCloudClient)
	if err != nil {
		log.WithFields(log.Fields{"package": "main","function": "DestroyVMs",}).Debugf("Unable to get VM actions. Virtual machine name: %s! Fatal Error!", vmName)
		log.WithFields(log.Fields{"package": "main","function": "DestroyVMs",}).Fatal(err)
	}

	error := vmactions.ExecuteDecommissionAction(catalogResID, oneCloudClient)
	if  error != nil {
		log.WithFields(log.Fields{"package": "main","function": "DestroyVMs",}).Debugf("Unable to get execute Decommision operation on the VM. Virtual machine name: %s! Fatal Error!", vmName)
		log.WithFields(log.Fields{"package": "main","function": "DestroyVMs",}).Fatal(error)
	}

	log.WithFields(log.Fields{"package": "main","function": "DestroyVMs",}).Debugf("Successfully destroyed VM: %s", vmName)
}