package gopowerstore

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	storageContainerMockURL = APIMockURL + storageContainerURL
)

var scID = "435669ba-28f5-4395-b5ca-6a7455726eaa"
var scID2 = "3765da74-28a7-49db-a693-10cec1de91f8"

func TestClientIMPL_CreateStorageContainer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, scID)
	httpmock.RegisterResponder("POST", storageContainerMockURL,
		httpmock.NewStringResponder(201, respData))

	createReq := StorageContainer{
		Name: "sc-test",
	}

	resp, err := C.CreateStorageContainer(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, scID, resp.ID)
}

func TestClientIMPL_DeleteStorageContainer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", storageContainerMockURL, scID),
		httpmock.NewStringResponder(204, ""))

	resp, err := C.DeleteStorageContainer(context.Background(), scID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_GetStorageContainer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, scID)
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", storageContainerMockURL, scID),
		httpmock.NewStringResponder(200, respData))
	storageContainer, err := C.GetStorageContainer(context.Background(), scID)
	assert.Nil(t, err)
	assert.Equal(t, scID, storageContainer.ID)
}
