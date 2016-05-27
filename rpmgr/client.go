package rpmgr

import (
	"net/url"

	"github.com/aocole/go-swagger-test/client"
	"github.com/aocole/go-swagger-test/client/iaas_resources"
	httpKitClient "github.com/go-openapi/runtime/client"
)

// UploadIAASResources upload the iaas resource to rpmgr.
func UploadIAASResources(targetURL *url.URL, iaasResources string) error {
	transport := httpKitClient.New(targetURL.Host, "/", []string{targetURL.Scheme})

	rpmgrClient := client.New(transport, nil)

	params := iaas_resources.NewUploadIAASResourcesParams().WithIaasResources(&iaasResources)
	err := rpmgrClient.IaasResources.UploadIAASResources(params)
	return err
}
