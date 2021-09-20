package gopowerstore

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	policyMockURL             = APIMockURL + policyURL
	replicationRuleMockURL    = APIMockURL + replicationRuleURL
	replicationSessionMockURL = APIMockURL + replicationSessionURL
)

func TestClientIMPL_CreateProtectionPolicy(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volID)
	httpmock.RegisterResponder("POST", policyMockURL,
		httpmock.NewStringResponder(201, respData))

	createReq := ProtectionPolicyCreate{
		Name:               "pp-test",
		Description:        "pp-test",
		ReplicationRuleIds: []string{"id"},
	}

	resp, err := C.CreateProtectionPolicy(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, volID, resp.ID)
}

func TestClientIMPL_CreateReplicationRule(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`{"id": "%s"}`, volID)
	httpmock.RegisterResponder("POST", replicationRuleMockURL,
		httpmock.NewStringResponder(201, respData))

	createReq := ReplicationRuleCreate{
		Name:           "rr-test",
		Rpo:            RpoFiveMinutes,
		RemoteSystemID: "XX-0000X",
	}

	resp, err := C.CreateReplicationRule(context.Background(), &createReq)
	assert.Nil(t, err)
	assert.Equal(t, volID, resp.ID)
}

func TestClientIMPL_DeleteProtectionPolicy(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", policyMockURL, volID),
		httpmock.NewStringResponder(204, ""))

	resp, err := C.DeleteProtectionPolicy(context.Background(), volID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_DeleteReplicationRule(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("DELETE", fmt.Sprintf("%s/%s", replicationRuleMockURL, volID),
		httpmock.NewStringResponder(204, ""))

	resp, err := C.DeleteReplicationRule(context.Background(), volID)
	assert.Nil(t, err)
	assert.Len(t, string(resp), 0)
}

func TestClientIMPL_GetProtectionPolicyByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", policyMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, volID)
	setResponder(respData)
	vol, err := C.GetProtectionPolicyByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, volID, vol.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetProtectionPolicyByName(context.Background(), "test")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.NotFound())
}

func TestClientIMPL_GetReplicationRuleByName(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setResponder := func(respData string) {
		httpmock.RegisterResponder("GET", replicationRuleMockURL,
			httpmock.NewStringResponder(200, respData))
	}
	respData := fmt.Sprintf(`[{"id": "%s"}]`, volID)
	setResponder(respData)
	replicationRule, err := C.GetReplicationRuleByName(context.Background(), "test")
	assert.Nil(t, err)
	assert.Equal(t, volID, replicationRule.ID)
	httpmock.Reset()
	setResponder("")
	_, err = C.GetReplicationRuleByName(context.Background(), "test")
	assert.NotNil(t, err)
	apiError := err.(APIError)
	assert.True(t, apiError.NotFound())
}

func TestClientIMPL_GetReplicationSessionByLocalResourceID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{
		"id":"%s",
		"name":"%s"
	}]`, volID, "session")
	httpmock.RegisterResponder("GET", replicationSessionMockURL,
		httpmock.NewStringResponder(200, respData))

	resp, err := C.GetReplicationSessionByLocalResourceID(context.Background(), volID2)
	assert.Nil(t, err)
	assert.Equal(t, volID, resp.ID)
}
