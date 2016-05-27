package rpmgr

import (
	"net/url"

	httpKitClient "github.com/go-openapi/runtime/client"
	"github.com/hpcloud/test/client"
	"github.com/hpcloud/test/client/iaas_resources"
)

// UploadIAASResources upload the iaas resource to rpmgr.
func UploadIAASResources(targetURL *url.URL, iaasResources string) error {
	transport := httpKitClient.New(targetURL.Host, "/", []string{targetURL.Scheme})

	rpmgrClient := client.New(transport, nil)

	params := iaas_resources.NewUploadIAASResourcesParams().WithIaasResources(&iaasResources)
	_, err := rpmgrClient.IaasResources.UploadIAASResources(params)
	return err
}
