package apiclient

import (
	"github.com/dghubble/sling"
	"log"
	"fmt"
	"net/http"
	"crypto/tls"
)

type OnecloudAPICLient struct {

	Username         string
	OneCloudHostName string
	OneCloudPort     string
	Tenant           string
	AuthToken        string
	HTTCloudNativelient  *sling.Sling
}


func NewAPIClient (username string, passwd string, tenant string, vraHost string, vraPort string ) (*OnecloudAPICLient,error) {

	transport := http.DefaultTransport.(*http.Transport)
	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	httCloudNativelient := sling.New().Base("https://"+vraHost+":"+vraPort+"/").
		          Set("Accept", "application/json").
		          Set("Content-Type", "application/json")
    authReq := &OneCloudAuthRequest{
    	        Username: username,
		        Password: passwd,
		        Tenant: tenant,
	}

	authRes := new(OneCloudAuthResponse)
	apiError := new(APIError)
	_,err := httCloudNativelient.Post("/identity/api/tokens").BodyJSON(authReq).
		     Receive(authRes,apiError)
	if err != nil {
		return nil,err
	}
	if !(len(apiError.Errors) == 0){
		log.Printf("%s\n", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
		return nil, fmt.Errorf("%s", fmt.Sprintf("vRealize API: %+v", apiError.Errors))
	}

	return &OnecloudAPICLient{
		Username: username,
		OneCloudHostName: vraHost,
		OneCloudPort: vraPort,
		Tenant: tenant,
		AuthToken: authRes.ID,
		HTTCloudNativelient: sling.New().Base("https://"+vraHost+":"+vraPort+"/").
			Set("Accept", "application/json").
			Set("Content-Type", "application/json").
			Set("Authorization", fmt.Sprintf("Bearer %s", authRes.ID)),
	}, nil
}