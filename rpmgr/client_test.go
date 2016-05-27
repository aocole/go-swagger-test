package rpmgr

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestUploadIAASResourcesSucceeds(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// w.Header().Add("Content-Type", "text/plain")
			// w.WriteHeader(http.StatusOK)
			return
		}))

	testURL, _ := url.Parse(server.URL)
	UploadIAASResources(testURL, "")
}
