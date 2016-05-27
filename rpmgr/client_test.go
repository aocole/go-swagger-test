package rpmgr

import (
	"net/url"
	"testing"
)

func TestUploadIAASResourcesSucceeds(t *testing.T) {
	testURL, _ := url.Parse("http://example.com")
	UploadIAASResources(testURL, "")
}
