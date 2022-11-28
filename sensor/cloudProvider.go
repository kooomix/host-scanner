package sensor

import (
	"net/http"

	"github.com/armosec/utils-go/httputils"
)

// CloudProviderInfo holds information about the Cloud Provider
type CloudProviderInfo struct {
	// Has access to cloud provider meta data API
	ProviderMetaDataAPIAccess bool `json:"providerMetaDataAPIAccess,omitempty"`
}

// SenseCloudProviderInfo returns `CloudProviderInfo`
func SenseCloudProviderInfo() (*CloudProviderInfo, error) {

	ret := CloudProviderInfo{}

	ret.ProviderMetaDataAPIAccess = hasMetaDataAPIAccess()

	return &ret, nil
}

// hasMetaDataAPIAccess - checks if there is an access to cloud provider meta data
var MetaDataAPIRequests = []struct {
	url     string
	headers map[string]string
} 	{
	{
		"http://169.254.169.254/computeMetadata/v1/?alt=json&recursive=true",
		map[string]string{"Metadata-Flavor": "Google"},
	},
	{
		"http://169.254.169.254/metadata/instance?api-version=2021-02-01",
		map[string]string{"Metadata": "true"},
	},
	{
		"http://169.254.169.254/latest/meta-data/local-hostname",
		map[string]string{},
	},
}

func hasMetaDataAPIAccess() bool {
	client := &http.Client{}

	for _,req := range MetaDataAPIRequests {
		res, err := httputils.HttpGet(client, req.url, req.headers)

		if err == nil && res.StatusCode == 200 {
			return true
		}	
	}

	return false

}
