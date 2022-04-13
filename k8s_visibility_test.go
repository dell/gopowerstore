package gopowerstore

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	k8sClusterMockURL = APIMockURL + k8sClusterURL
)

var k8sClusterID = "5e8d7b7b-671b-336f-db4e-cee0fbdc981e"

func TestClientIMPL_RegisterK8sCluster(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, k8sClusterID)
	httpmock.RegisterResponder("POST", k8sClusterMockURL,
		httpmock.NewStringResponder(201, respData))
	registerReq := K8sCluster{
		Name:      "test-cluster",
		IPAddress: "1.1.1.1",
		Port:      8080,
		Token:     "test-token",
	}

	cl, err := C.RegisterK8sCluster(context.Background(), &registerReq)
	assert.Nil(t, err)
	assert.Equal(t, k8sClusterID, cl.ID)
}
